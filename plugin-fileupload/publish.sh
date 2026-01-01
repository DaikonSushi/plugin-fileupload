#!/bin/bash

# Plugin FileUpload - Quick Start Script
# This script helps you quickly publish the plugin to GitHub

set -e

echo "🚀 Plugin FileUpload - Quick Start"
echo "=================================="
echo ""

# Check if we're in the right directory
if [ ! -f "main.go" ]; then
    echo "❌ Error: main.go not found. Please run this script from the plugin-fileupload directory."
    exit 1
fi

# Get GitHub username
read -p "Enter your GitHub username (default: DaikonSushi): " GITHUB_USER
GITHUB_USER=${GITHUB_USER:-DaikonSushi}

# Get repository name
read -p "Enter repository name (default: plugin-fileupload): " REPO_NAME
REPO_NAME=${REPO_NAME:-plugin-fileupload}

# Get version
read -p "Enter version tag (default: v1.0.0): " VERSION
VERSION=${VERSION:-v1.0.0}

echo ""
echo "📋 Configuration:"
echo "   GitHub User: $GITHUB_USER"
echo "   Repository: $REPO_NAME"
echo "   Version: $VERSION"
echo ""

read -p "Continue? (y/n): " CONFIRM
if [ "$CONFIRM" != "y" ]; then
    echo "❌ Aborted."
    exit 0
fi

echo ""
echo "🔧 Step 1: Checking Git status..."

if [ -d ".git" ]; then
    echo "✅ Git repository already initialized"
else
    echo "📦 Initializing Git repository..."
    git init
    echo "✅ Git repository initialized"
fi

echo ""
echo "🔧 Step 2: Adding files..."
git add .

echo ""
echo "🔧 Step 3: Committing changes..."
if git diff-index --quiet HEAD --; then
    echo "ℹ️  No changes to commit"
else
    git commit -m "Initial commit: File upload test plugin for bot-platform"
    echo "✅ Changes committed"
fi

echo ""
echo "🔧 Step 4: Setting up remote..."
REMOTE_URL="https://github.com/$GITHUB_USER/$REPO_NAME.git"

if git remote | grep -q "origin"; then
    echo "ℹ️  Remote 'origin' already exists"
    CURRENT_URL=$(git remote get-url origin)
    if [ "$CURRENT_URL" != "$REMOTE_URL" ]; then
        echo "⚠️  Current remote URL: $CURRENT_URL"
        echo "⚠️  Expected URL: $REMOTE_URL"
        read -p "Update remote URL? (y/n): " UPDATE_REMOTE
        if [ "$UPDATE_REMOTE" = "y" ]; then
            git remote set-url origin "$REMOTE_URL"
            echo "✅ Remote URL updated"
        fi
    fi
else
    git remote add origin "$REMOTE_URL"
    echo "✅ Remote 'origin' added"
fi

echo ""
echo "🔧 Step 5: Setting main branch..."
git branch -M main

echo ""
echo "🔧 Step 6: Pushing to GitHub..."
echo "ℹ️  You may need to enter your GitHub credentials"
git push -u origin main

echo ""
echo "🔧 Step 7: Creating and pushing tag..."
if git tag | grep -q "$VERSION"; then
    echo "⚠️  Tag $VERSION already exists"
    read -p "Delete and recreate tag? (y/n): " RECREATE_TAG
    if [ "$RECREATE_TAG" = "y" ]; then
        git tag -d "$VERSION"
        git push origin :refs/tags/"$VERSION" 2>/dev/null || true
        git tag -a "$VERSION" -m "Release $VERSION: File upload test plugin"
        git push origin "$VERSION"
        echo "✅ Tag recreated and pushed"
    fi
else
    git tag -a "$VERSION" -m "Release $VERSION: File upload test plugin"
    git push origin "$VERSION"
    echo "✅ Tag created and pushed"
fi

echo ""
echo "✅ Done! Plugin published to GitHub"
echo ""
echo "📦 Next steps:"
echo "   1. Visit: https://github.com/$GITHUB_USER/$REPO_NAME"
echo "   2. Check Actions tab for build status"
echo "   3. Wait for build to complete (~3-5 minutes)"
echo "   4. Check Releases tab for binaries"
echo ""
echo "🔧 Install in bot-platform:"
echo "   cd /path/to/bot-platform"
echo "   ./botctl install https://github.com/$GITHUB_USER/$REPO_NAME"
echo "   ./botctl start fileupload"
echo ""
echo "🧪 Test commands:"
echo "   /filehelp"
echo "   /testfile"
echo "   /testfile medium json"
echo "   /createfile test.txt Hello World!"
echo ""
