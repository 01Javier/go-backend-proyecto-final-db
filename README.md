## Correr el proyecto
crear un archivo `.env` y configurar los parametros especificados en `.env.example`
```bash
go mod tidy
go run server/main.go
```

Por si acaso, ejecuta: `[Environment]::SetEnvironmentVariable("CGO_ENABLED", "1", "User")` (powershell)

---

## Problema 1: Errores de compilación con godror

### Mensaje de error

```
# github.com/godror/godror
undefined: VersionInfo
undefined: StartupMode
undefined: ShutdownMode
too many errors
```

### Causa

Posiblemente falte `Oracle Instant Client` descargar [aqui](https://www.oracle.com/database/technologies/instant-client/winx64-64-downloads.html)

luego copiar la carpeta a "C:\Program Files\Common Files\Oracle\instantclient_23_9" despues agregarla al PATH


Si no funciona, probar esto: 

**Paso 1: Limpia el cache de modulos**

```powershell
go clean -modcache
```

**Paso 2: Usar una version anterior de godror**

```powershell
go get github.com/godror/godror@v0.40.3
go mod tidy
```


**Otras versiones estables para probar:**

- `v0.40.3` (recomendada)
- `v0.33.3`
- `v0.44.0` (última, pero puede tener problemas)

---

## Problema 2: Compilador GCC no encontrado

### Mensaje de error

```
# runtime/cgo
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
```

### Causa

`godror` utiliza CGO (enlaces en C), lo que requiere un compilador de C en Windows.

### Solución

**Opción A: Instalar con Chocolatey (más fácil)**

```powershell
# Instala Chocolatey si no lo tienes
Set-ExecutionPolicy Bypass -Scope Process -Force
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# Instala MinGW
choco install mingw -y

# Refresca el entorno
refreshenv
```

**Opción B: Instalación manual de MinGW**

1. Descarga MinGW-w64:
   - URL: <https://github.com/niXman/mingw-builds-binaries/releases>
   - Archivo: `x86_64-*-release-win32-seh-msvcrt-*.7z`

2. Extrae en `C:\mingw64`

3. Agrega a la variable PATH:

```powershell
$env:Path += ";C:\mingw64\bin"

[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\mingw64\bin", "User")
```

4. Reiniciar PowerShell y verificar:

```powershell
gcc --version
```

**Opción C: Instalar MSYS2**

1. Descarga desde: <https://www.msys2.org/>
2. Instala y abre la terminal de MSYS2
3. Ejecuta:

```bash
pacman -S mingw-w64-x86_64-gcc
```

4. Agrega a la variable PATH: `C:\msys64\mingw64\bin`

**Después de instalar GCC:**

```powershell
# Habilita CGO
$env:CGO_ENABLED = "1"

# Configura de forma permanente
[Environment]::SetEnvironmentVariable("CGO_ENABLED", "1", "User")

# Prueba
go run server/main.go
```

---
