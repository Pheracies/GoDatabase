---
name: git-windows-repair
description: Provides steps to diagnose and repair Windows Git process locks (.git/index.lock) and corrupted Git object/HEAD references.
---

# Windows Git Process Lock & Repository Repair Skill

Use this skill when Git commands on Windows fail with:
- `fatal: Unable to create '.git/index.lock': File exists`
- `fatal: could not parse HEAD`
- `unable to open loose object ...: Function not implemented`

## Steps

### 1. Process & Lock Diagnostics
Check for hanging `git.exe` background tasks holding file locks:
```powershell
Get-Process git -ErrorAction SilentlyContinue | Stop-Process -Force
```

### 2. Remove File Locks & Hidden Attributes
On Windows, `.git` files can have `Hidden` or `ReadOnly` attributes preventing deletion. Clear attributes before removing:
```powershell
cmd /c "attrib -r -h -s .git\index.lock"
Remove-Item -Force .git/index.lock -ErrorAction SilentlyContinue
```

### 3. Rebuilding Index / Corrupt Refs
If Git HEAD or object references are damaged (`fatal: could not parse HEAD`):
```powershell
# Clear damaged index
Remove-Item -Force .git/index -ErrorAction SilentlyContinue
git add .
```

If `.git/HEAD` or ref files are corrupt:
```powershell
cmd /c "attrib -h .git"
git init -b main
git add .
```
