# Plugin FileUpload - Quick Reference

## 🚀 Quick Start

```bash
# 1. Test locally
cd plugin-fileupload
./test.sh

# 2. Publish to GitHub
./publish.sh

# 3. Install in bot-platform
cd ../bot-platform
./botctl install https://github.com/DaikonSushi/plugin-fileupload
./botctl start fileupload
```

## 📝 Commands

| Command | Description | Example |
|---------|-------------|---------|
| `/filehelp` | Show help | `/filehelp` |
| `/testfile [size] [type]` | Create test file | `/testfile medium json` |
| `/createfile <name> <content>` | Create custom file | `/createfile test.txt Hello!` |
| `/uploadgroup <path> [name] [folder]` | Upload to group | `/uploadgroup /tmp/test.txt` |
| `/uploadprivate <path> [name]` | Upload to private | `/uploadprivate /tmp/test.txt` |

## 📊 File Sizes

- **small**: 10 lines
- **medium**: 100 lines
- **large**: 1000 lines

## 📄 File Types

- **txt**: Plain text
- **json**: JSON format
- **md**: Markdown format

## 🧪 Test Scenarios

### Scenario 1: Quick Test
```
/testfile
```
Creates and uploads a small text file.

### Scenario 2: Different Sizes
```
/testfile small txt
/testfile medium json
/testfile large md
```

### Scenario 3: Custom Content
```
/createfile hello.txt Hello from bot-platform!
/createfile config.json {"debug":true}
```

### Scenario 4: Group Upload with Folder
```
/uploadgroup /tmp/report.txt monthly_report.txt /reports
```

### Scenario 5: Private Upload
```
/uploadprivate /tmp/data.json user_data.json
```

## 🔧 Management

```bash
# Start plugin
./botctl start fileupload

# Stop plugin
./botctl stop fileupload

# Restart plugin
./botctl restart fileupload

# Check status
./botctl status fileupload

# View logs
./botctl logs fileupload
```

## 📦 File Structure

```
plugin-fileupload/
├── .github/
│   └── workflows/
│       └── release.yml      # CI/CD configuration
├── .gitignore               # Git ignore rules
├── DEPLOYMENT.md            # Deployment guide
├── LICENSE                  # MIT License
├── README.md                # Main documentation
├── QUICKREF.md              # This file
├── go.mod                   # Go module file
├── main.go                  # Plugin source code
├── publish.sh               # Publish script
└── test.sh                  # Test script
```

## 🐛 Troubleshooting

### Plugin won't start
```bash
./botctl status fileupload
./botctl logs fileupload
chmod +x plugins-bin/fileupload-plugin
```

### Upload fails
- Check NapCat connection
- Verify file path exists
- Check file permissions
- Review bot-platform logs

### Build fails on GitHub
- Ensure no `replace` directive in go.mod
- Check GitHub Actions logs
- Verify Go version compatibility

## 📚 Resources

- [Full README](README.md)
- [Deployment Guide](DEPLOYMENT.md)
- [bot-platform](https://github.com/DaikonSushi/bot-platform)
- [NapCat](https://napcat.napneko.icu/)

## 💡 Tips

1. **Test locally first**: Run `./test.sh` before publishing
2. **Use descriptive filenames**: Makes it easier to identify files
3. **Check logs**: Use `./botctl logs fileupload` for debugging
4. **Start small**: Test with small files first
5. **Monitor uploads**: Watch for success/error messages

## 🎯 Success Checklist

- [ ] Plugin builds successfully locally
- [ ] `./test.sh` passes all tests
- [ ] Published to GitHub
- [ ] GitHub Actions build completes
- [ ] Release created with binaries
- [ ] Installed in bot-platform
- [ ] Plugin starts without errors
- [ ] `/filehelp` works
- [ ] `/testfile` uploads successfully
- [ ] Files appear in QQ

## 📞 Support

For issues or questions:
1. Check logs: `./botctl logs fileupload`
2. Review [DEPLOYMENT.md](DEPLOYMENT.md)
3. Create GitHub Issue
4. Check bot-platform documentation
