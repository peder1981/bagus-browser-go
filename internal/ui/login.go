package ui

import (
	"log"
	"os"
	"path/filepath"

	"github.com/peder1981/bagus-browser-go/internal/config"
	"github.com/peder1981/bagus-browser-go/internal/security"
	webview "github.com/webview/webview_go"
)

// LoginDialog representa o diálogo de login
type LoginDialog struct {
	w        webview.WebView
	username string
	userPath string
}

// NewLoginDialog cria novo diálogo de login
func NewLoginDialog() *LoginDialog {
	return &LoginDialog{}
}

// Show exibe o diálogo de login
func (l *LoginDialog) Show() (string, error) {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()

	l.w = w

	w.SetTitle("Bagus Browser - Login")
	w.SetSize(600, 400, webview.HintNone)

	// Bind funções
	w.Bind("validateUsername", func(username string) string {
		if err := security.ValidateUsername(username); err != nil {
			return err.Error()
		}
		return ""
	})

	w.Bind("startBrowser", func(username string) {
		l.startBrowser(username)
	})

	// HTML do login
	html := l.generateHTML()
	w.SetHtml(html)

	w.Run()

	return l.userPath, nil
}

// startBrowser inicia o browser com username
func (l *LoginDialog) startBrowser(username string) bool {
	// Valida username
	if err := security.ValidateUsername(username); err != nil {
		log.Printf("Username inválido: %v", err)
		return false
	}

	// Define path do usuário
	userPath := filepath.Join("/tmp", username)

	// Verifica se diretório existe, se não, cria
	if _, err := os.Stat(userPath); os.IsNotExist(err) {
		if err := os.MkdirAll(userPath, 0700); err != nil {
			log.Printf("Erro ao criar diretório: %v", err)
			return false
		}
	}

	// Cria config.json se não existir
	configPath := filepath.Join(userPath, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		cfg := config.Default()
		if err := cfg.Save(configPath); err != nil {
			log.Printf("Erro ao criar configuração: %v", err)
			return false
		}
	}

	// Cria subdiretórios
	dirs := []string{"log", "analyze", "analyze/pending", "default"}
	for _, dir := range dirs {
		path := filepath.Join(userPath, dir)
		if err := os.MkdirAll(path, 0700); err != nil {
			log.Printf("Erro ao criar diretório %s: %v", dir, err)
			return false
		}
	}

	l.username = username
	l.userPath = userPath

	// Fecha janela de login
	l.w.Terminate()

	return true
}

// generateHTML gera HTML do formulário de login
func (l *LoginDialog) generateHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Bagus Browser - Login</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
        }
        .container {
            background: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 10px 40px rgba(0,0,0,0.2);
            width: 400px;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 28px;
        }
        .subtitle {
            color: #666;
            margin-bottom: 30px;
            font-size: 14px;
        }
        .form-group {
            margin-bottom: 20px;
        }
        label {
            display: block;
            margin-bottom: 8px;
            color: #333;
            font-weight: 500;
        }
        input {
            width: 100%;
            padding: 12px;
            border: 2px solid #e0e0e0;
            border-radius: 6px;
            font-size: 14px;
            transition: border-color 0.3s;
        }
        input:focus {
            outline: none;
            border-color: #667eea;
        }
        .error {
            color: #e74c3c;
            font-size: 12px;
            margin-top: 5px;
            display: none;
        }
        .info {
            background: #f8f9fa;
            padding: 15px;
            border-radius: 6px;
            margin-bottom: 20px;
            font-size: 13px;
            color: #555;
        }
        .info ul {
            margin: 10px 0 0 20px;
        }
        .info li {
            margin: 5px 0;
        }
        button {
            width: 100%;
            padding: 14px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 6px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s;
        }
        button:hover {
            transform: translateY(-2px);
        }
        button:active {
            transform: translateY(0);
        }
        .script-info {
            margin-top: 20px;
            padding: 15px;
            background: #f8f9fa;
            border-radius: 6px;
            font-size: 12px;
            color: #666;
            display: none;
        }
        .script-info code {
            display: block;
            background: #2d2d2d;
            color: #f8f8f2;
            padding: 10px;
            border-radius: 4px;
            margin-top: 10px;
            font-family: 'Courier New', monospace;
            overflow-x: auto;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Bagus Browser</h1>
        <p class="subtitle">Browser Seguro e Privado</p>
        
        <div class="info">
            <strong>Requisitos do Username:</strong>
            <ul>
                <li>Entre 3 e 32 caracteres</li>
                <li>Apenas letras, números, _ ou -</li>
                <li>Sem espaços ou caracteres especiais</li>
            </ul>
        </div>

        <div class="form-group">
            <label for="username">Username:</label>
            <input type="text" id="username" placeholder="Digite seu username" 
                   oninput="validateInput()" onkeydown="if(event.key==='Enter') startBrowserClick()">
            <div class="error" id="error"></div>
        </div>

        <button onclick="startBrowserClick()">Iniciar Browser</button>

        <div class="script-info" id="scriptInfo">
            <strong>Diretório criado:</strong>
            <code id="pathInfo"></code>
        </div>
    </div>

    <script>
        function validateInput() {
            const username = document.getElementById('username').value;
            const errorDiv = document.getElementById('error');
            
            if (username.length === 0) {
                errorDiv.style.display = 'none';
                return;
            }

            // Validação client-side simples
            if (username.length < 3 || username.length > 32) {
                errorDiv.textContent = 'Username deve ter entre 3 e 32 caracteres';
                errorDiv.style.display = 'block';
                return;
            }

            if (!/^[a-zA-Z0-9_-]+$/.test(username)) {
                errorDiv.textContent = 'Username deve conter apenas letras, números, _ ou -';
                errorDiv.style.display = 'block';
                return;
            }

            errorDiv.style.display = 'none';
        }

        function startBrowserClick() {
            const username = document.getElementById('username').value.trim();
            const errorDiv = document.getElementById('error');
            
            if (!username) {
                errorDiv.textContent = 'Por favor, digite um username';
                errorDiv.style.display = 'block';
                return;
            }

            // Validação client-side
            if (username.length < 3 || username.length > 32) {
                errorDiv.textContent = 'Username deve ter entre 3 e 32 caracteres';
                errorDiv.style.display = 'block';
                return;
            }

            if (!/^[a-zA-Z0-9_-]+$/.test(username)) {
                errorDiv.textContent = 'Username deve conter apenas letras, números, _ ou -';
                errorDiv.style.display = 'block';
                return;
            }

            // Limpa erro e inicia browser
            errorDiv.style.display = 'none';
            startBrowser(username);
        }

        // Foco automático no input
        window.addEventListener('load', () => {
            document.getElementById('username').focus();
        });
    </script>
</body>
</html>`
}
