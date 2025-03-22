import os
import sys

try:
    from rich import console
except ImportError:
    os.system("pip install rich")

console = console.Console()


def run_cmd(cmd: str):
    os.system(cmd)
    console.log(f"执行命令：{cmd}", style="bold green")


def main():
    if len(sys.argv) != 2:
        console.log("Usage: python make.py build|run|clean")
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
    elif sys.argv[1] == "clean":
        run_cmd("go clean -cache")
        os.chdir("./cli")
        run_cmd("go mod tidy")
        os.chdir("../")
        os.chdir("./web")
        run_cmd("go mod tidy")
        os.chdir("../")

    else:
        console.log("Usage: python make.py build|run|clean")


def linux_run():
    linux_build()
    console.log("Linux 运行")
    run_cmd("bin/finix-web")
    pass


def windows_run():
    windows_build()
    console.log("Windows 运行")
    run_cmd(r".\bin\finix-web.exe")
    pass


def linux_build():
    console.log("Linux 编译")
    os.chdir("./cli")
    run_cmd("go build -x -o ../bin/finix")
    run_cmd("chmod +x ../bin/finix")
    os.chdir("../")
    console.log("已生成 bin/finix")
    os.chdir("./web")
    run_cmd("go build -x -o ../bin/finix-web")
    run_cmd("chmod +x ../bin/finix-web")
    os.chdir("../")
    console.log("已生成 bin/finix-web")
    console.log("Linux 编译完成")
    pass


def windows_build():
    console.log("Windows 编译")
    os.chdir("./cli")
    run_cmd("go build -x -o ../bin/finix.exe")
    os.chdir("../")
    console.log("已生成 bin/finix.exe")
    os.chdir("./web")
    run_cmd("go build -x -o ../bin/finix-web.exe")
    os.chdir("../")
    console.log("已生成 bin/finix-web.exe")
    console.log("Windows 编译完成")
    pass


if __name__ == "__main__":
    main()
