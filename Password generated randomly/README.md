这是用GO语言写的一个小工具，在本地运行的随机密码生成器。下载这个文件夹，直接运行.exe文件。如果为了美观一点，可以添加快捷方式。
修改快捷方式的图标与名字，可以用文件夹里的图片或者你喜欢的图片。如果要本地编译，Windows要下载go语言环境，需要用到粘贴板，
我们需要引入一个可靠的第三方库，因为 Go 标准库本身不包含剪贴板操作。要下载我们将使用 github.com/atotto/clipboard 这个库。
go mod init password-gen
go get github.com/atotto/clipboard
