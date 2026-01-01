// plugin-fileupload - A comprehensive file upload test plugin for bot-platform
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DaikonSushi/bot-platform/pkg/pluginsdk"
)

// FileUploadPlugin tests file upload functionality
type FileUploadPlugin struct {
	bot *pluginsdk.BotClient
}

// Info returns plugin metadata
func (p *FileUploadPlugin) Info() pluginsdk.PluginInfo {
	return pluginsdk.PluginInfo{
		Name:              "fileupload",
		Version:           "1.0.0",
		Description:       "Comprehensive file upload test plugin for NapCat integration",
		Author:            "hovanzhang",
		Commands:          []string{"testfile", "uploadgroup", "uploadprivate", "createfile", "filehelp"},
		HandleAllMessages: false,
	}
}

// OnStart is called when the plugin starts
func (p *FileUploadPlugin) OnStart(bot *pluginsdk.BotClient) error {
	p.bot = bot
	bot.Log("info", "File upload test plugin started!")
	return nil
}

// OnStop is called when the plugin stops
func (p *FileUploadPlugin) OnStop() error {
	return nil
}

// OnMessage handles incoming messages
func (p *FileUploadPlugin) OnMessage(ctx context.Context, bot *pluginsdk.BotClient, msg *pluginsdk.Message) bool {
	return false
}

// OnCommand handles commands
func (p *FileUploadPlugin) OnCommand(ctx context.Context, bot *pluginsdk.BotClient, cmd string, args []string, msg *pluginsdk.Message) bool {
	switch cmd {
	case "filehelp":
		p.handleHelp(bot, msg)
		return true
	case "testfile":
		p.handleTestFile(bot, args, msg)
		return true
	case "createfile":
		p.handleCreateFile(bot, args, msg)
		return true
	case "uploadgroup":
		p.handleUploadGroup(bot, args, msg)
		return true
	case "uploadprivate":
		p.handleUploadPrivate(bot, args, msg)
		return true
	}
	return false
}

// handleHelp shows plugin help information
func (p *FileUploadPlugin) handleHelp(bot *pluginsdk.BotClient, msg *pluginsdk.Message) {
	bot.Reply(msg,
		pluginsdk.Text("📤 File Upload Test Plugin\n"),
		pluginsdk.Text("━━━━━━━━━━━━━━━━━━━━\n"),
		pluginsdk.Text("Available Commands:\n\n"),
		pluginsdk.Text("📝 /testfile [size] [type]\n"),
		pluginsdk.Text("   Create and upload a test file\n"),
		pluginsdk.Text("   size: small/medium/large (default: small)\n"),
		pluginsdk.Text("   type: txt/json/md (default: txt)\n\n"),
		pluginsdk.Text("📄 /createfile <name> <content>\n"),
		pluginsdk.Text("   Create a custom file with content\n\n"),
		pluginsdk.Text("📤 /uploadgroup <path> [name] [folder]\n"),
		pluginsdk.Text("   Upload file to group\n\n"),
		pluginsdk.Text("📨 /uploadprivate <path> [name]\n"),
		pluginsdk.Text("   Upload file to private chat\n\n"),
		pluginsdk.Text("❓ /filehelp\n"),
		pluginsdk.Text("   Show this help message"),
	)
}

// handleTestFile creates and uploads a test file
func (p *FileUploadPlugin) handleTestFile(bot *pluginsdk.BotClient, args []string, msg *pluginsdk.Message) {
	// Parse arguments
	size := "small"
	fileType := "txt"
	
	if len(args) > 0 {
		size = strings.ToLower(args[0])
	}
	if len(args) > 1 {
		fileType = strings.ToLower(args[1])
	}

	// Validate size
	if size != "small" && size != "medium" && size != "large" {
		bot.Reply(msg, pluginsdk.Text("❌ Invalid size. Use: small, medium, or large"))
		return
	}

	// Validate file type
	if fileType != "txt" && fileType != "json" && fileType != "md" {
		bot.Reply(msg, pluginsdk.Text("❌ Invalid type. Use: txt, json, or md"))
		return
	}

	// Create test file
	tmpDir := os.TempDir()
	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("test_%s_%s.%s", size, timestamp, fileType)
	testFile := filepath.Join(tmpDir, fileName)

	content := p.generateTestContent(size, fileType)
	
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Failed to create test file: %v", err)))
		return
	}
	defer os.Remove(testFile)

	// Get file info
	fileInfo, _ := os.Stat(testFile)
	fileSize := fileInfo.Size()

	bot.Reply(msg, pluginsdk.Text(fmt.Sprintf(
		"✅ Test file created:\n"+
			"📄 Name: %s\n"+
			"📊 Size: %d bytes\n"+
			"📝 Type: %s\n"+
			"⏱️  Time: %s\n\n"+
			"📤 Uploading...",
		fileName, fileSize, fileType, timestamp,
	)))

	// Upload based on message type
	if msg.GroupID > 0 {
		err = bot.UploadGroupFile(msg.GroupID, testFile, fileName)
		if err != nil {
			bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Group upload failed: %v", err)))
		} else {
			bot.Reply(msg, pluginsdk.Text("✅ File uploaded to group successfully!"))
		}
	} else {
		err = bot.UploadPrivateFile(msg.UserID, testFile, fileName)
		if err != nil {
			bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Private upload failed: %v", err)))
		} else {
			bot.Reply(msg, pluginsdk.Text("✅ File uploaded to private chat successfully!"))
		}
	}
}

// generateTestContent generates test content based on size and type
func (p *FileUploadPlugin) generateTestContent(size, fileType string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	
	var lines int
	switch size {
	case "small":
		lines = 10
	case "medium":
		lines = 100
	case "large":
		lines = 1000
	}

	switch fileType {
	case "json":
		content := "{\n"
		content += fmt.Sprintf("  \"plugin\": \"fileupload\",\n")
		content += fmt.Sprintf("  \"timestamp\": \"%s\",\n", timestamp)
		content += fmt.Sprintf("  \"size\": \"%s\",\n", size)
		content += "  \"data\": [\n"
		for i := 0; i < lines; i++ {
			comma := ","
			if i == lines-1 {
				comma = ""
			}
			content += fmt.Sprintf("    {\"line\": %d, \"content\": \"Test data line %d\"}%s\n", i+1, i+1, comma)
		}
		content += "  ]\n}\n"
		return content

	case "md":
		content := "# File Upload Test\n\n"
		content += fmt.Sprintf("**Plugin:** fileupload\n")
		content += fmt.Sprintf("**Timestamp:** %s\n", timestamp)
		content += fmt.Sprintf("**Size:** %s\n\n", size)
		content += "## Test Data\n\n"
		for i := 0; i < lines; i++ {
			content += fmt.Sprintf("- Line %d: Test data content for file upload testing\n", i+1)
		}
		return content

	default: // txt
		content := fmt.Sprintf("File Upload Test - %s\n", timestamp)
		content += strings.Repeat("=", 50) + "\n\n"
		content += fmt.Sprintf("Plugin: fileupload\n")
		content += fmt.Sprintf("Size: %s (%d lines)\n", size, lines)
		content += fmt.Sprintf("Type: %s\n\n", fileType)
		content += "Test Data:\n"
		content += strings.Repeat("-", 50) + "\n"
		for i := 0; i < lines; i++ {
			content += fmt.Sprintf("Line %04d: This is test data for file upload functionality testing\n", i+1)
		}
		return content
	}
}

// handleCreateFile creates a custom file with user-provided content
func (p *FileUploadPlugin) handleCreateFile(bot *pluginsdk.BotClient, args []string, msg *pluginsdk.Message) {
	if len(args) < 2 {
		bot.Reply(msg,
			pluginsdk.Text("📄 Create Custom File\n"),
			pluginsdk.Text("━━━━━━━━━━━━━━\n"),
			pluginsdk.Text("Usage:\n"),
			pluginsdk.Text("  /createfile <filename> <content>\n\n"),
			pluginsdk.Text("Example:\n"),
			pluginsdk.Text("  /createfile test.txt Hello World!\n"),
			pluginsdk.Text("  /createfile data.json {\"key\":\"value\"}"),
		)
		return
	}

	fileName := args[0]
	content := strings.Join(args[1:], " ")

	tmpDir := os.TempDir()
	testFile := filepath.Join(tmpDir, fileName)

	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Failed to create file: %v", err)))
		return
	}
	defer os.Remove(testFile)

	fileInfo, _ := os.Stat(testFile)
	bot.Reply(msg, pluginsdk.Text(fmt.Sprintf(
		"✅ File created: %s (%d bytes)\n📤 Uploading...",
		fileName, fileInfo.Size(),
	)))

	// Upload based on message type
	if msg.GroupID > 0 {
		err = bot.UploadGroupFile(msg.GroupID, testFile, fileName)
		if err != nil {
			bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Upload failed: %v", err)))
		} else {
			bot.Reply(msg, pluginsdk.Text("✅ File uploaded successfully!"))
		}
	} else {
		err = bot.UploadPrivateFile(msg.UserID, testFile, fileName)
		if err != nil {
			bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Upload failed: %v", err)))
		} else {
			bot.Reply(msg, pluginsdk.Text("✅ File uploaded successfully!"))
		}
	}
}

// handleUploadGroup uploads a file to group
func (p *FileUploadPlugin) handleUploadGroup(bot *pluginsdk.BotClient, args []string, msg *pluginsdk.Message) {
	if msg.GroupID == 0 {
		bot.Reply(msg, pluginsdk.Text("❌ This command can only be used in groups"))
		return
	}

	if len(args) < 1 {
		bot.Reply(msg,
			pluginsdk.Text("📤 Upload Group File\n"),
			pluginsdk.Text("━━━━━━━━━━━━━━\n"),
			pluginsdk.Text("Usage:\n"),
			pluginsdk.Text("  /uploadgroup <filepath> [name] [folder]\n\n"),
			pluginsdk.Text("Examples:\n"),
			pluginsdk.Text("  /uploadgroup /tmp/test.txt\n"),
			pluginsdk.Text("  /uploadgroup /tmp/test.txt myfile.txt\n"),
			pluginsdk.Text("  /uploadgroup /tmp/test.txt myfile.txt /docs"),
		)
		return
	}

	filePath := args[0]
	fileName := filepath.Base(filePath)
	if len(args) > 1 {
		fileName = args[1]
	}

	folder := "/"
	if len(args) > 2 {
		folder = args[2]
	}

	// Check if file exists
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ File not found: %s", filePath)))
		return
	}

	bot.Reply(msg, pluginsdk.Text(fmt.Sprintf(
		"📤 Uploading to group:\n"+
			"📄 File: %s\n"+
			"📊 Size: %d bytes\n"+
			"📝 Name: %s\n"+
			"📁 Folder: %s",
		filePath, fileInfo.Size(), fileName, folder,
	)))

	err = bot.UploadGroupFile(msg.GroupID, filePath, fileName, folder)
	if err != nil {
		bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Upload failed: %v", err)))
	} else {
		bot.Reply(msg, pluginsdk.Text("✅ File uploaded to group successfully!"))
	}
}

// handleUploadPrivate uploads a file to private chat
func (p *FileUploadPlugin) handleUploadPrivate(bot *pluginsdk.BotClient, args []string, msg *pluginsdk.Message) {
	if msg.GroupID > 0 {
		bot.Reply(msg, pluginsdk.Text("❌ This command can only be used in private chats"))
		return
	}

	if len(args) < 1 {
		bot.Reply(msg,
			pluginsdk.Text("📨 Upload Private File\n"),
			pluginsdk.Text("━━━━━━━━━━━━━━\n"),
			pluginsdk.Text("Usage:\n"),
			pluginsdk.Text("  /uploadprivate <filepath> [name]\n\n"),
			pluginsdk.Text("Examples:\n"),
			pluginsdk.Text("  /uploadprivate /tmp/test.txt\n"),
			pluginsdk.Text("  /uploadprivate /tmp/test.txt myfile.txt"),
		)
		return
	}

	filePath := args[0]
	fileName := filepath.Base(filePath)
	if len(args) > 1 {
		fileName = args[1]
	}

	// Check if file exists
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ File not found: %s", filePath)))
		return
	}

	bot.Reply(msg, pluginsdk.Text(fmt.Sprintf(
		"📨 Uploading to private chat:\n"+
			"📄 File: %s\n"+
			"📊 Size: %d bytes\n"+
			"📝 Name: %s",
		filePath, fileInfo.Size(), fileName,
	)))

	err = bot.UploadPrivateFile(msg.UserID, filePath, fileName)
	if err != nil {
		bot.Reply(msg, pluginsdk.Text(fmt.Sprintf("❌ Upload failed: %v", err)))
	} else {
		bot.Reply(msg, pluginsdk.Text("✅ File uploaded to private chat successfully!"))
	}
}

func main() {
	pluginsdk.Run(&FileUploadPlugin{})
}
