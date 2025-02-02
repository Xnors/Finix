import os
import sys

try:
    from rich import console
except ImportError:
    os.system("pip install rich")

console = console.Console()

def main():
    if len(sys.argv) != 2:
        console.log("Usage: python make.py build|run")
        return

    if not os.path.exists("bin"):
        os.mkdir("bin")

    # 命令行参数
    if sys.argv[1] == "build":
        if os.name == "posix":
            linux_build()
        elif os.name == "nt":
            windows_build()
    elif sys.argv[1] == "run":
        if os.name == "posix":
            linux_run()
        elif os.name == "nt":
            windows_run()
    else:
        console.log("Usage: python make.py build|run")


def linux_run():
    linux_build()
    console.log("Linux 运行")
    os.system("bin/finix-web")
    pass


def windows_run():
    windows_build()
    console.log("Windows 运行")
    os.system("start http://localhost:8080")
    os.system(r".\bin\finix-web.exe")
    pass


def linux_build():
    console.log("Linux 编译")
    os.chdir("./cli")
    os.system("go build -x -o ../bin/finix")
    os.system("chmod +x ../bin/finix")
    os.chdir("../")
    console.log("已生成 bin/finix")
    os.chdir("./web")
    os.system("go build -x -o ../bin/finix-web")
    os.system("chmod +x ../bin/finix-web")
    os.chdir("../")
    console.log("已生成 bin/finix-web")
    console.log("Linux 编译完成")
    pass


def windows_build():
    console.log("Windows 编译")
    os.chdir("./cli")
    os.system("go build -x -o ../bin/finix.exe")
    os.chdir("../")
    console.log("已生成 bin/finix.exe")
    os.chdir("./web")
    os.system("go build -x -o ../bin/finix-web.exe")
    os.chdir("../")
    console.log("已生成 bin/finix-web.exe")
    console.log("Windows 编译完成")
    pass


if __name__ == "__main__":
    main()
