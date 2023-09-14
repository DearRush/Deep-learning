package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func preCheck() bool {
	// 检查文件夹是否存在：6-E3_data文件夹是否存在
	if _, err := os.Stat("6-E3_data"); os.IsNotExist(err) {
		return false
	}

	// 删除文件夹：6-E3_data_processed
	// 如果存在，则删除
	if _, err := os.Stat("6-E3_data_processed"); !os.IsNotExist(err) {
		os.RemoveAll("6-E3_data_processed")
	}

	// 检查文件夹是否存在：6-E3_data_processed
	// 如果不存在，则创建文件夹
	if _, err := os.Stat("6-E3_data_processed"); os.IsNotExist(err) {
		os.Mkdir("6-E3_data_processed", os.ModePerm)
	}

	// 检查文件夹是否存在：6-E3_data_processed/test_img
	// 如果不存在，则创建文件夹
	if _, err := os.Stat("6-E3_data_processed/test_img"); os.IsNotExist(err) {
		os.Mkdir("6-E3_data_processed/test_img", os.ModePerm)
	}

	// 检查文件夹是否存在：6-E3_data_processed/train_img
	// 如果不存在，则创建文件夹
	if _, err := os.Stat("6-E3_data_processed/train_img"); os.IsNotExist(err) {
		os.Mkdir("6-E3_data_processed/train_img", os.ModePerm)
	}

	// 检查文件夹是否存在：6-E3_data_processed/val_img
	// 如果不存在，则创建文件夹
	if _, err := os.Stat("6-E3_data_processed/val_img"); os.IsNotExist(err) {
		os.Mkdir("6-E3_data_processed/val_img", os.ModePerm)
	}

	// 检查文件是否存在：6-E3_data/Test_labels.txt
	// 如果不存在，则报错
	if _, err := os.Stat("6-E3_data/Test_labels.txt"); os.IsNotExist(err) {
		return false
	}

	// 检查文件是否存在：6-E3_data/train_labels.txt
	// 如果不存在，则报错
	if _, err := os.Stat("6-E3_data/train_labels.txt"); os.IsNotExist(err) {
		return false
	}

	// 检查文件是否存在：6-E3_data/val_labels.txt
	// 如果不存在，则报错
	if _, err := os.Stat("6-E3_data/val_labels.txt"); os.IsNotExist(err) {
		return false
	}

	return true
}

func main() {
	// 执行预检查
	if !preCheck() {
		panic("preCheck failed, please check the files and folders, put data in the right place.")
	}

	fmt.Println("preCheck success, start to process data.")

	if !process_testImage() {
		panic("process_testImage failed.")
	} else {
		fmt.Println("process_testImage success.")
	}

	if !process_trainImg() {
		panic("process_trainImg failed.")
	} else {
		fmt.Println("process_trainImg success.")
	}

	if !process_valImg() {
		panic("process_valImg failed.")
	} else {
		fmt.Println("process_valImg success.")
	}

	fmt.Println("process data success.")
	fmt.Println("please check the folder: 6-E3_data_processed")
}

func process_testImage() bool {
	// 读取文件：6-E3_data/Test_labels.txt
	// 打开文件
	file, err := os.Open("6-E3_data/Test_labels.txt")
	if err != nil {
		panic(err)
	}

	// 关闭文件
	defer file.Close()

	// 读取文件内容
	// 逐行读取文件
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// 解析行数据
		fields := strings.Split(line, "\t")
		if len(fields) == 2 {
			filename := fields[0]
			folderStr := fields[1]

			// 转换文件夹数字为整数
			folder, err := strconv.Atoi(folderStr)
			if err != nil {
				fmt.Println("无法解析文件夹数字:", err)
				continue
			}

			// 检查文件夹"6-E3_data_processed/test_img/"+strconv.Itoa(folder)是否存在
			foldTarget := "6-E3_data_processed/test_img/" + strconv.Itoa(folder)

			// 如果不存在，则创建文件夹
			if _, err := os.Stat(foldTarget); os.IsNotExist(err) {
				os.Mkdir(foldTarget, os.ModePerm)
			}

			// 处理文件和文件夹
			fmt.Printf("文件名：%s\t文件夹：%d\n", filename, folder)
			// 在这里你可以编写代码来复制文件到对应的文件夹
			// 例如：6-E3_data_processed/test_img/1/1.jpg 6-E3_data_processed/test_img/2/2.jpg
			copyFile("6-E3_data/test_img/"+filename, "6-E3_data_processed/test_img/"+strconv.Itoa(folder)+"/"+filename)
		} else {
			fmt.Println("无效的行数据:", line)
			panic("无效的行数据")
		}
	}

	return true
}

func process_trainImg() bool {
	// 读取文件：6-E3_data/train_labels.txt
	// 打开文件
	file, err := os.Open("6-E3_data/train_labels.txt")
	if err != nil {
		panic(err)
	}

	// 关闭文件
	defer file.Close()

	// 读取文件内容
	// 逐行读取文件
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// 解析行数据
		fields := strings.Split(line, "\t")
		if len(fields) == 2 {
			filename := fields[0]
			folderStr := fields[1]

			// 转换文件夹数字为整数
			folder, err := strconv.Atoi(folderStr)
			if err != nil {
				fmt.Println("无法解析文件夹数字:", err)
				continue
			}

			// 检查文件夹"6-E3_data_processed/train_img/"+strconv.Itoa(folder)是否存在
			foldTarget := "6-E3_data_processed/train_img/" + strconv.Itoa(folder)

			// 如果不存在，则创建文件夹
			if _, err := os.Stat(foldTarget); os.IsNotExist(err) {
				os.Mkdir(foldTarget, os.ModePerm)
			}

			// 处理文件和文件夹
			fmt.Printf("文件名：%s\t文件夹：%d\n", filename, folder)
			// 在这里你可以编写代码来复制文件到对应的文件夹
			// 例如：6-E3_data_processed/train_img/1/1.jpg 6-E3_data_processed/train_img/2/2.jpg
			copyFile("6-E3_data/train_img/"+filename, "6-E3_data_processed/train_img/"+strconv.Itoa(folder)+"/"+filename)
		} else {
			fmt.Println("无效的行数据:", line)
			panic("无效的行数据")
		}
	}

	return true	
}

func process_valImg() bool {
	// 读取文件：6-E3_data/val_labels.txt
	// 打开文件
	file, err := os.Open("6-E3_data/val_labels.txt")
	if err != nil {
		panic(err)
	}

	// 关闭文件
	defer file.Close()

	// 读取文件内容
	// 逐行读取文件
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// 解析行数据
		fields := strings.Split(line, "\t")
		if len(fields) == 2 {
			filename := fields[0]
			folderStr := fields[1]

			// 转换文件夹数字为整数
			folder, err := strconv.Atoi(folderStr)
			if err != nil {
				fmt.Println("无法解析文件夹数字:", err)
				continue
			}

			// 检查文件夹"6-E3_data_processed/val_img/"+strconv.Itoa(folder)是否存在
			foldTarget := "6-E3_data_processed/val_img/" + strconv.Itoa(folder)

			// 如果不存在，则创建文件夹
			if _, err := os.Stat(foldTarget); os.IsNotExist(err) {
				os.Mkdir(foldTarget, os.ModePerm)
			}

			// 处理文件和文件夹
			fmt.Printf("文件名：%s\t文件夹：%d\n", filename, folder)
			// 在这里你可以编写代码来复制文件到对应的文件夹
			// 例如：6-E3_data_processed/val_img/1/1.jpg 6-E3_data_processed/val_img/2/2.jpg
			copyFile("6-E3_data/val_img/"+filename, "6-E3_data_processed/val_img/"+strconv.Itoa(folder)+"/"+filename)
		} else {
			fmt.Println("无效的行数据:", line)
			panic("无效的行数据")
		}
	}

	return true
}

func copyFile(src, dest string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("无法打开源文件: %s", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("无法创建目标文件: %s", err)
	}
	defer destFile.Close()

	// 拷贝文件内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("拷贝文件时出错: %s", err)
	}

	return nil
}


