# Plugin FileUpload

一个综合性的文件上传测试插件，用于测试 bot-platform 通过 NapCat 上传文件的功能。

## 功能特性

### 📝 自动测试文件生成
- `/testfile [size] [type]` - 创建并上传测试文件
  - **size**: `small`(10行) / `medium`(100行) / `large`(1000行)
  - **type**: `txt` / `json` / `md`
  - 示例: `/testfile medium json`

### 📄 自定义文件创建
- `/createfile <filename> <content>` - 创建自定义内容的文件
  - 示例: `/createfile test.txt Hello World!`
  - 示例: `/createfile data.json {"key":"value"}`

### 📤 群组文件上传
- `/uploadgroup <filepath> [name] [folder]` - 上传文件到群组
  - 示例: `/uploadgroup /tmp/test.txt`
  - 示例: `/uploadgroup /tmp/test.txt myfile.txt /documents`

### 📨 私聊文件上传
- `/uploadprivate <filepath> [name]` - 上传文件到私聊
  - 示例: `/uploadprivate /tmp/test.txt`
  - 示例: `/uploadprivate /tmp/test.txt myfile.txt`

### ❓ 帮助信息
- `/filehelp` - 显示插件帮助信息

## 本地编译测试

### 本地开发环境

在本地开发时，需要使用本地的 bot-platform 代码。确保 `go.mod` 中的 replace 指令已启用：

```go
replace github.com/DaikonSushi/bot-platform => /Users/hovanzhang/git_repo/napcat/bot-platform
```

### 编译步骤

```bash
# 更新依赖
go mod tidy

# 在本地编译
go build -ldflags="-s -w" -o fileupload-plugin .

# 测试插件信息输出
./fileupload-plugin --info
```

**注意**: 发布到 GitHub 时，GitHub Actions 会自动移除 replace 指令并使用正确的依赖版本。

## 发布到 GitHub

### 1. 创建 GitHub 仓库

在 GitHub 上创建新仓库 `plugin-fileupload`

### 2. 初始化并推送代码

```bash
cd plugin-fileupload
git init
git add .
git commit -m "Initial commit: File upload test plugin"
git branch -M main
git remote add origin https://github.com/DaikonSushi/plugin-fileupload.git
git push -u origin main
```

### 3. 创建 Release 触发自动构建

#### 方法 1: 使用发布脚本（推荐）

```bash
# 使用自动化脚本发布
./publish.sh 1.0.0
```

脚本会自动：
- ✅ 检查版本号
- ✅ 验证 git 状态
- ✅ 测试本地编译
- ✅ 创建并推送 tag
- ✅ 触发 GitHub Actions

#### 方法 2: 手动发布

```bash
# 创建标签
git tag v1.0.0
git push origin v1.0.0
```

GitHub Actions 会自动构建多平台二进制文件并创建 Release：
- `fileupload-plugin_linux_amd64`
- `fileupload-plugin_linux_arm64`
- `fileupload-plugin_darwin_amd64`
- `fileupload-plugin_darwin_arm64`
- `fileupload-plugin_windows_amd64.exe`

## 在 bot-platform 中使用

### 安装插件

```bash
# 从 GitHub Release 安装
./botctl install https://github.com/DaikonSushi/plugin-fileupload

# 或手动下载并放置
cp fileupload-plugin plugins-bin/
```

### 配置插件

创建 `plugins-config/fileupload.json`:

```json
{
  "name": "fileupload",
  "enabled": true,
  "binary": "fileupload-plugin",
  "description": "File upload test plugin"
}
```

### 管理插件

```bash
# 启动插件
./botctl start fileupload

# 停止插件
./botctl stop fileupload

# 查看状态
./botctl status fileupload

# 重启插件
./botctl restart fileupload
```

## 测试场景

### 场景 1: 快速测试
在群组或私聊中发送:
```
/testfile
```
将创建并上传一个小型文本文件

### 场景 2: 测试不同大小
```
/testfile small txt
/testfile medium json
/testfile large md
```

### 场景 3: 自定义内容
```
/createfile hello.txt Hello from bot-platform!
/createfile config.json {"debug":true,"port":8080}
```

### 场景 4: 指定文件夹（群组）
```
/uploadgroup /tmp/test.txt report.txt /reports
```

## 技术细节

- **语言**: Go 1.24
- **依赖**: bot-platform SDK
- **通信**: gRPC
- **支持平台**: Linux, macOS, Windows (amd64/arm64)

## 开发说明

本地开发时，取消 `go.mod` 中的 replace 注释：

```go
replace github.com/DaikonSushi/bot-platform => /path/to/local/bot-platform
```

发布前记得注释掉该行，让 CI 使用远程依赖。

## License

MIT

## 作者

hovanzhang
