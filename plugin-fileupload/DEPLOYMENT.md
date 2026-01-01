# Plugin FileUpload 部署指南

本指南将帮助你将 plugin-fileupload 发布到 GitHub 并测试文件上传功能。

## 前置准备

1. **GitHub 账号**: 确保你有 GitHub 账号
2. **Git 配置**: 本地已配置 Git
3. **bot-platform**: 已部署并运行 bot-platform
4. **NapCat**: 已配置 NapCat 连接

## 步骤 1: 创建 GitHub 仓库

1. 访问 https://github.com/new
2. 仓库名称: `plugin-fileupload`
3. 描述: `File upload test plugin for bot-platform with NapCat`
4. 选择 Public
5. 不要初始化 README (我们已经有了)
6. 点击 "Create repository"

## 步骤 2: 推送代码到 GitHub

```bash
cd /Users/hovanzhang/git_repo/napcat/plugin-fileupload

# 初始化 Git 仓库
git init

# 添加所有文件
git add .

# 提交
git commit -m "Initial commit: File upload test plugin for bot-platform"

# 设置主分支
git branch -M main

# 添加远程仓库 (替换为你的 GitHub 用户名)
git remote add origin https://github.com/DaikonSushi/plugin-fileupload.git

# 推送代码
git push -u origin main
```

## 步骤 3: 创建 Release 触发构建

```bash
# 创建版本标签
git tag -a v1.0.0 -m "Release v1.0.0: Initial release with comprehensive file upload testing"

# 推送标签到 GitHub
git push origin v1.0.0
```

## 步骤 4: 等待 GitHub Actions 构建

1. 访问你的仓库页面
2. 点击 "Actions" 标签
3. 查看 "Build and Release" workflow 运行状态
4. 等待所有构建完成 (约 3-5 分钟)

构建完成后会生成以下文件:
- `fileupload-plugin_linux_amd64`
- `fileupload-plugin_linux_arm64`
- `fileupload-plugin_darwin_amd64`
- `fileupload-plugin_darwin_arm64`
- `fileupload-plugin_windows_amd64.exe`

## 步骤 5: 验证 Release

1. 访问仓库的 "Releases" 页面
2. 应该能看到 `v1.0.0` release
3. 确认所有平台的二进制文件都已上传

## 步骤 6: 在 bot-platform 中安装插件

### 方法 1: 使用 botctl 安装 (推荐)

```bash
cd /Users/hovanzhang/git_repo/napcat/bot-platform

# 从 GitHub 安装
./botctl install https://github.com/DaikonSushi/plugin-fileupload

# 启动插件
./botctl start fileupload
```

### 方法 2: 手动安装

```bash
cd /Users/hovanzhang/git_repo/napcat/bot-platform

# 下载对应平台的二进制文件
# macOS ARM64 示例:
curl -L -o plugins-bin/fileupload-plugin \
  https://github.com/DaikonSushi/plugin-fileupload/releases/download/v1.0.0/fileupload-plugin_darwin_arm64

# 添加执行权限
chmod +x plugins-bin/fileupload-plugin

# 创建配置文件
cat > plugins-config/fileupload.json <<EOF
{
  "name": "fileupload",
  "enabled": true,
  "binary": "fileupload-plugin",
  "description": "File upload test plugin"
}
EOF

# 启动插件
./botctl start fileupload
```

## 步骤 7: 测试文件上传功能

### 测试 1: 基础测试文件上传

在 QQ 群组或私聊中发送:
```
/testfile
```

预期结果:
- 插件创建一个小型文本测试文件
- 文件通过 NapCat 上传到 QQ
- 收到成功确认消息

### 测试 2: 不同大小和类型

```
/testfile small txt
/testfile medium json
/testfile large md
```

预期结果:
- 创建不同大小和格式的测试文件
- 所有文件都能成功上传

### 测试 3: 自定义文件内容

```
/createfile hello.txt Hello from bot-platform!
/createfile data.json {"test":true,"value":123}
```

预期结果:
- 创建包含自定义内容的文件
- 文件成功上传

### 测试 4: 群组文件夹上传 (仅群组)

```
/uploadgroup /tmp/test.txt report.txt /reports
```

预期结果:
- 文件上传到群组的指定文件夹

### 测试 5: 查看帮助

```
/filehelp
```

预期结果:
- 显示完整的命令帮助信息

## 步骤 8: 查看日志

```bash
cd /Users/hovanzhang/git_repo/napcat/bot-platform

# 查看插件日志
./botctl logs fileupload

# 查看 bot-platform 日志
tail -f bot-platform.log
```

## 故障排查

### 问题 1: 插件无法启动

```bash
# 检查插件状态
./botctl status fileupload

# 查看详细日志
./botctl logs fileupload

# 检查二进制文件权限
ls -la plugins-bin/fileupload-plugin
chmod +x plugins-bin/fileupload-plugin
```

### 问题 2: 文件上传失败

检查:
1. NapCat 是否正常运行
2. bot-platform 与 NapCat 的连接状态
3. 文件路径是否正确
4. 文件权限是否正确

```bash
# 测试 NapCat 连接
curl http://localhost:3000/health

# 查看 bot-platform 日志
tail -f bot-platform.log
```

### 问题 3: GitHub Actions 构建失败

1. 检查 go.mod 中是否有 replace 指令 (应该被注释掉)
2. 检查 GitHub Actions 日志
3. 确认 Go 版本兼容性

## 验证清单

- [ ] GitHub 仓库创建成功
- [ ] 代码推送到 GitHub
- [ ] Release v1.0.0 创建成功
- [ ] 所有平台二进制文件构建成功
- [ ] 插件在 bot-platform 中安装成功
- [ ] 插件启动成功
- [ ] `/testfile` 命令测试通过
- [ ] `/createfile` 命令测试通过
- [ ] `/uploadgroup` 命令测试通过 (群组)
- [ ] `/uploadprivate` 命令测试通过 (私聊)
- [ ] `/filehelp` 显示正常
- [ ] 文件成功通过 NapCat 上传到 QQ

## 后续更新

当需要发布新版本时:

```bash
cd /Users/hovanzhang/git_repo/napcat/plugin-fileupload

# 修改代码...

# 提交更改
git add .
git commit -m "Update: description of changes"
git push

# 创建新版本标签
git tag -a v1.0.1 -m "Release v1.0.1: bug fixes and improvements"
git push origin v1.0.1
```

GitHub Actions 会自动构建并创建新的 Release。

## 参考资料

- [bot-platform 文档](https://github.com/DaikonSushi/bot-platform)
- [NapCat 文档](https://napcat.napneko.icu/)
- [GitHub Actions 文档](https://docs.github.com/en/actions)

## 支持

如有问题，请在 GitHub 仓库创建 Issue。
