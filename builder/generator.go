package builder

import (
	"context"

	"cshop-website/utils"
	"cshop-website/view/landingpage"
	"cshop-website/view/privacypage"
	"cshop-website/view/termspage"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func RunStaticGenerator() {
	fmt.Println("🏗️  Starting Static Site Generation...")
	outputDir := "dist"
	langs := []string{"en", "ar"}

	// 1. Clear and recreate dist folder
	os.RemoveAll(outputDir)
	os.MkdirAll(outputDir, 0755)

	// 2. Generate the root index.html redirect
	generateRootRedirect(outputDir)

	// 3. Copy assets folder
	fmt.Println("📁 Copying assets...")
	err := copyDir("assets", filepath.Join(outputDir, "assets"))
	if err != nil {
		fmt.Printf("❌ Failed to copy assets: %v\n", err)
	}

	// 4. Generate language-specific pages
	for _, lang := range langs {
		fmt.Printf("🌐 Processing Language: %s\n", lang)

		basePath := filepath.Join(outputDir, lang)
		os.MkdirAll(basePath, 0755)
		os.MkdirAll(filepath.Join(basePath, "privacy"), 0755)
		os.MkdirAll(filepath.Join(basePath, "terms"), 0755)

		saveHTML(filepath.Join(basePath, "index.html"), lang, "index")
		saveHTML(filepath.Join(basePath, "privacy", "index.html"), lang, "privacy")
		saveHTML(filepath.Join(basePath, "terms", "index.html"), lang, "terms")
	}

	fmt.Println("✅ Generation complete! Created in /dist folder.")
}

func saveHTML(targetPath string, lang string, pageType string) {
	f, err := os.Create(targetPath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", targetPath, err)
	}
	defer f.Close()

	content := utils.LandingPage[lang]
	privacyPolicy := utils.PrivacyContent[lang]
	termsOfUse := utils.TermsContent[lang]
	ctx := context.Background()

	switch pageType {
	case "index":
		landingpage.Show(lang, content).Render(ctx, f)
	case "privacy":
		privacypage.Show(lang, content, privacyPolicy).Render(ctx, f)
	case "terms":
		termspage.Show(lang, content, termsOfUse).Render(ctx, f)
	}
}

// generateRootRedirect creates the dist/index.html file
func generateRootRedirect(outputDir string) {
	filePath := filepath.Join(outputDir, "index.html")
	html := `<!DOCTYPE html>
<html lang="ar" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="refresh" content="0; url=./ar/">
    <link rel="canonical" href="./ar/">
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
	<link rel="alternate" hreflang="en" href="https://classic-shop.ly/en" />
    <link rel="alternate" hreflang="ar" href="https://classic-shop.ly/ar" />
    <link rel="alternate" hreflang="x-default" href="https://classic-shop.ly/ar" />
	<link rel="icon" type="image/png" href="/assets/png/logo.png" />
    <title>Classic</title>
</head>
</html>`

	_ = os.WriteFile(filePath, []byte(html), 0644)
}

// copyDir recursively copies a directory tree
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(src, path)
		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		return copyFile(path, targetPath)
	})
}

// copyFile is a helper for copyDir to handle individual files
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
