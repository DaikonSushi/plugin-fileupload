#!/bin/bash

# Plugin FileUpload - Local Test Script
# This script helps you test the plugin locally before publishing

set -e

echo "🧪 Plugin FileUpload - Local Test"
echo "================================="
echo ""

# Check if we're in the right directory
if [ ! -f "main.go" ]; then
    echo "❌ Error: main.go not found. Please run this script from the plugin-fileupload directory."
    exit 1
fi

echo "🔧 Step 1: Building plugin..."
go build -o fileupload-plugin .
echo "✅ Build successful"

echo ""
echo "🔧 Step 2: Testing plugin info..."
./fileupload-plugin --info
echo "✅ Plugin info test passed"

echo ""
echo "🔧 Step 3: Checking binary size..."
SIZE=$(ls -lh fileupload-plugin | awk '{print $5}')
echo "📦 Binary size: $SIZE"

echo ""
echo "🔧 Step 4: Testing file creation..."
TEST_DIR=$(mktemp -d)
echo "📁 Test directory: $TEST_DIR"

# Create test files
echo "Creating test files..."
echo "Test content" > "$TEST_DIR/test.txt"
echo '{"test": true}' > "$TEST_DIR/test.json"
echo "# Test" > "$TEST_DIR/test.md"

echo "✅ Test files created:"
ls -lh "$TEST_DIR"

echo ""
echo "🔧 Step 5: Verifying file operations..."
if [ -f "$TEST_DIR/test.txt" ] && [ -f "$TEST_DIR/test.json" ] && [ -f "$TEST_DIR/test.md" ]; then
    echo "✅ All test files verified"
else
    echo "❌ Some test files missing"
    exit 1
fi

echo ""
echo "🧹 Cleaning up test files..."
rm -rf "$TEST_DIR"
echo "✅ Cleanup complete"

echo ""
echo "✅ All local tests passed!"
echo ""
echo "📋 Plugin Information:"
echo "   Name: fileupload"
echo "   Version: 1.0.0"
echo "   Binary: fileupload-plugin"
echo ""
echo "🚀 Next steps:"
echo "   1. Run './publish.sh' to publish to GitHub"
echo "   2. Or manually test with bot-platform:"
echo "      - Copy fileupload-plugin to bot-platform/plugins-bin/"
echo "      - Create config in bot-platform/plugins-config/fileupload.json"
echo "      - Start with: ./botctl start fileupload"
echo ""
echo "🧪 Test commands in QQ:"
echo "   /filehelp - Show help"
echo "   /testfile - Quick test"
echo "   /testfile medium json - Test with options"
echo "   /createfile hello.txt Hello World! - Custom file"
echo ""
