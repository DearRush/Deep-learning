import torch
from os.path import exists
import torch.nn as nn
import torch.nn.functional as F
from torch.utils.data import DataLoader
from torch.optim import SGD
import torch.optim as optim
import torchvision.transforms
import matplotlib.pyplot as plt
from define import CNN

transformer=torchvision.transforms.Compose([torchvision.transforms.ToTensor()])
batchsize=100
train_loader = DataLoader(torchvision.datasets.ImageFolder('/lustre/home/acct-stu/stu202/protein_identify/6data/train_img',transform=transformer),batch_size=batchsize,shuffle=True)
val_loader=DataLoader(torchvision.datasets.ImageFolder('/lustre/home/acct-stu/stu202/protein_identify/6data/val_img',transform=transformer),batch_size=batchsize,shuffle=True)
test_loader = DataLoader(torchvision.datasets.ImageFolder('/lustre/home/acct-stu/stu202/protein_identify/6data/test_img',transform=transformer),batch_size=batchsize,shuffle=True)


testimg=enumerate(test_loader)
#index, (imagedata,label)= next(images)
#print(imagedata)
#print(imagedata.shape)

e_poch=1
all_loss=0
all_number=0
true_number=0
best_accuracy=0
y_loss=[30]
x_iter=[0]
y_accuracy=[0]
x_epoch=[0]
net=CNN()
optimizer=SGD(net.parameters(),lr=0.005,momentum=0.9,weight_decay=0.0005)


while e_poch<=6:
    images = enumerate(train_loader)
    valimg=enumerate(val_loader)
    net.train()
    for index,(imagedata,label) in images:
        optimizer.zero_grad()
        output=net.foward(imagedata)
        loss=F.cross_entropy(output,label)
        loss.backward()
        optimizer.step()
        if (index*batchsize)%1000==0 and index>=1:
            all_loss+=loss.item()
            y_loss.append(all_loss)
            x_iter.append(index*batchsize+65000*(e_poch-1))
            print("for epoch %d %d/%d loss is %.2f"%(e_poch,index*batchsize,65000,all_loss))
            all_loss=0
        else:
            all_loss+=loss.item()
    all_loss=0
    net.eval()
    for index,(imagedata,label) in valimg:
        output=net.foward(imagedata)
        all_number+=batchsize
        output=output.detach()
        pred = output.data.max(1, keepdim=True)[1]
        true_number += pred.eq(label.data.view_as(pred)).sum()
    accuracy=(float)(true_number/all_number)
    true_number=0
    all_number=0
    if accuracy>best_accuracy:
        best_accuracy=accuracy
        torch.save(net.state_dict(),"./model.pth")
        torch.save(optimizer.state_dict(), './optimizer.pth')
    print("for this epoch%d, the accuracy of validation is %.2f"%(e_poch,accuracy))
    y_accuracy.append(accuracy)
    x_epoch.append(e_poch)
    e_poch+=1

if exists("model.pth"):
    network_state_dict = torch.load('model.pth')
    net.load_state_dict(network_state_dict)
    optimizer_state_dict = torch.load('optimizer.pth')
    optimizer.load_state_dict(optimizer_state_dict) 
net.eval()
all_number=0
true_number=0
for index,(imagedata,label) in testimg:
        output=net.foward(imagedata)
        all_number+=batchsize
        output=output.detach()
        pred = output.data.max(1, keepdim=True)[1]
        true_number += pred.eq(label.data.view_as(pred)).sum()
accuracy=(float)(true_number/all_number)
print("For test set, the overall accuracy is %.2f"%(accuracy))

plt.plot(x_iter,y_loss,color="blue")
plt.xlabel('Number of training images')
plt.ylabel('Loss')
plt.title("Loss function ver training images")
plt.savefig("Loss_cnn.png")

plt.plot(x_epoch,y_accuracy,color="red")
plt.xlabel('Epoch time')
plt.ylabel('Accuracy')
plt.title("Accuracy of each epoch")
plt.savefig("Accuracy.png")


