# minecraft-mod-updater

🚀 Minecraft Mod Pack Auto-Updater

A simple, single-file Go application that automatically downloads and installs the latest version of a Minecraft mod pack to the user's local mods folder.

This application simplifies mod pack distribution by checking a single remote URL for the latest mod pack ZIP file, ensuring all players have the same files without needing to delete or manage old versions.

🛠️ Setup Instructions for Your Own Mod Pack

Follow these steps to customize this updater for your specific mod pack:

Step 1: Host Your Mod Pack Files (GitHub Releases)

You need to host your complete mod pack as a single ZIP file and get a permanent download URL.

Create Your ZIP: Gather all your .jar mod files and compress them into a single file named something like modpack-v1.0.0.zip. The .jar files should be at the root of the archive (not inside a subfolder).

Create a Public Repository: If you haven't already, create a public GitHub repository (e.g., username/modpack-files).

Create a Release: Go to the Releases tab in your repository and click "Draft a new release." (If the Releases tab is missing, create a temporary README.md first).

Upload: Use the "Attach binaries" section to upload your modpack-v1.0.0.zip.

Publish: Publish the release (you must assign a tag like v1.0.0).

Copy the URL: Right-click the uploaded ZIP file link under Assets and copy the link address. This is your direct ZIP download link.

Step 2: Set Up the Permanent Manifest URL (Vercel/Web Host)

This is the permanent link the application checks for the latest mod pack URL.

Create latest_url.txt: In your website's public file directory (e.g., the public folder of your Vercel-hosted React site), create a plain text file named latest_url.txt.

Paste the URL: Inside this file, paste the direct ZIP download link you copied in Step 1.

Commit and Deploy: Commit and deploy your website to Vercel.

Final Manifest URL: Your permanent manifest URL will be your site's domain plus the file name (e.g., https://my-modpack.vercel.app/latest_url.txt).

Step 3: Update the Go Source Code

You only need to update the single hard-coded constant in the Go source code.

Open the downloader.go file.

Replace the placeholder URL for ManifestURL with your actual, permanent Vercel URL from Step 2.

// downloader.go

// This must be the permanent URL pointing to your 'latest_url.txt' file.
const ManifestURL = "[https://my-modpack.vercel.app/latest_url.txt](https://my-modpack.vercel.app/latest_url.txt)" 


Step 4: Compile and Distribute

Compile the application to create a dependency-free executable for your users.

Navigate to your project folder and run the appropriate command based on the target operating system (using PowerShell syntax):

Target OS

Command (in PowerShell)

Windows

$env:GOOS="windows"; $env:GOARCH="amd64"; go build -ldflags="-s -w" -o modupdater-windows.exe .

macOS

$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -ldflags="-s -w" -o modupdater-macos .

Linux

$env:GOOS="linux"; $env:GOARCH="amd64"; go build -ldflags="-s -w" -o modupdater-linux .

Distribute the resulting executable (e.g., modupdater-windows.exe) to your friends.

🔁 How to Update Your Mod Pack (No Code Changes Needed!)

Once the Go executable is distributed, future updates are simple:

Create your new ZIP file (e.g., modpack-v2.0.0.zip).

Create a new GitHub Release and upload modpack-v2.0.0.zip.

Copy the new direct download link for the v2.0.0 ZIP file.

Update latest_url.txt: Edit the latest_url.txt file on your web host (Vercel) and paste the new v2.0.0 URL inside it, replacing the old one.

Commit and push the latest_url.txt update to Vercel.

The distributed updater will now automatically check the permanent Vercel link and download the new v2.0.0 mod pack the next time it runs!