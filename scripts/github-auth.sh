#!/bin/bash
set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Configurações
TOKEN_CACHE_DIR="${HOME}/.config/bagus-browser"
TOKEN_CACHE_FILE="${TOKEN_CACHE_DIR}/github_token.cache"
TOKEN_CACHE_ENCRYPTED="${TOKEN_CACHE_FILE}.enc"

# Criar diretório se não existir
mkdir -p "${TOKEN_CACHE_DIR}"
chmod 700 "${TOKEN_CACHE_DIR}"

# Função para criptografar token
encrypt_token() {
    local token="$1"
    local key=$(hostname | sha256sum | cut -d' ' -f1)
    echo -n "$token" | openssl enc -aes-256-cbc -salt -pbkdf2 -pass pass:"$key" -out "${TOKEN_CACHE_ENCRYPTED}" 2>/dev/null
    chmod 600 "${TOKEN_CACHE_ENCRYPTED}"
}

# Função para descriptografar token
decrypt_token() {
    if [ ! -f "${TOKEN_CACHE_ENCRYPTED}" ]; then
        return 1
    fi
    
    local key=$(hostname | sha256sum | cut -d' ' -f1)
    local token=$(openssl enc -aes-256-cbc -d -pbkdf2 -pass pass:"$key" -in "${TOKEN_CACHE_ENCRYPTED}" 2>/dev/null)
    
    if [ $? -eq 0 ] && [ -n "$token" ]; then
        echo "$token"
        return 0
    fi
    return 1
}

# Função para validar token
validate_token() {
    local token="$1"
    
    if [ -z "$token" ]; then
        return 1
    fi
    
    # Testar token com API do GitHub
    local response=$(curl -s -o /dev/null -w "%{http_code}" \
        -H "Authorization: token ${token}" \
        -H "Accept: application/vnd.github.v3+json" \
        https://api.github.com/user)
    
    if [ "$response" = "200" ]; then
        return 0
    fi
    return 1
}

# Função para obter informações do usuário
get_user_info() {
    local token="$1"
    curl -s \
        -H "Authorization: token ${token}" \
        -H "Accept: application/vnd.github.v3+json" \
        https://api.github.com/user | jq -r '.login // "unknown"'
}

# Função principal para obter token
get_github_token() {
    # 1. Verificar variável de ambiente
    if [ -n "$GITHUB_TOKEN" ]; then
        echo -e "${BLUE}🔑 Usando GITHUB_TOKEN da variável de ambiente${NC}" >&2
        if validate_token "$GITHUB_TOKEN"; then
            echo -e "${GREEN}✅ Token válido${NC}" >&2
            echo "$GITHUB_TOKEN"
            return 0
        else
            echo -e "${YELLOW}⚠️  Token da variável de ambiente inválido${NC}" >&2
        fi
    fi
    
    # 2. Verificar cache
    echo -e "${BLUE}🔍 Verificando cache de token...${NC}" >&2
    local cached_token=$(decrypt_token)
    if [ -n "$cached_token" ]; then
        if validate_token "$cached_token"; then
            local username=$(get_user_info "$cached_token")
            echo -e "${GREEN}✅ Token em cache válido (usuário: ${username})${NC}" >&2
            echo "$cached_token"
            return 0
        else
            echo -e "${YELLOW}⚠️  Token em cache expirado ou inválido${NC}" >&2
            rm -f "${TOKEN_CACHE_ENCRYPTED}"
        fi
    fi
    
    # 3. Tentar gh CLI
    if command -v gh &> /dev/null; then
        echo -e "${BLUE}🔍 Verificando gh CLI...${NC}" >&2
        if gh auth status &> /dev/null; then
            local gh_token=$(gh auth token 2>/dev/null)
            if [ -n "$gh_token" ] && validate_token "$gh_token"; then
                local username=$(get_user_info "$gh_token")
                echo -e "${GREEN}✅ Token do gh CLI válido (usuário: ${username})${NC}" >&2
                
                # Cachear token do gh CLI
                encrypt_token "$gh_token"
                echo -e "${GREEN}💾 Token cacheado para uso futuro${NC}" >&2
                
                echo "$gh_token"
                return 0
            fi
        fi
    fi
    
    # 4. Solicitar novo token
    echo -e "${YELLOW}⚠️  Nenhum token válido encontrado${NC}" >&2
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}" >&2
    echo -e "${YELLOW}Por favor, configure um token do GitHub:${NC}" >&2
    echo "" >&2
    echo -e "${BLUE}Opção 1: Criar token${NC}" >&2
    echo -e "  1. Ir para: ${GREEN}https://github.com/settings/tokens${NC}" >&2
    echo -e "  2. Generate new token (classic)" >&2
    echo -e "  3. Marcar: ${GREEN}repo${NC} (Full control)" >&2
    echo -e "  4. Copiar token" >&2
    echo "" >&2
    echo -e "${BLUE}Opção 2: Usar gh CLI${NC}" >&2
    echo -e "  ${GREEN}gh auth login${NC}" >&2
    echo "" >&2
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}" >&2
    
    # Perguntar se quer inserir token agora
    read -p "Deseja inserir um token agora? (s/n): " -n 1 -r >&2
    echo "" >&2
    
    if [[ $REPLY =~ ^[Ss]$ ]]; then
        read -sp "Cole o token do GitHub: " new_token >&2
        echo "" >&2
        
        if validate_token "$new_token"; then
            local username=$(get_user_info "$new_token")
            echo -e "${GREEN}✅ Token válido! (usuário: ${username})${NC}" >&2
            
            # Cachear token
            encrypt_token "$new_token"
            echo -e "${GREEN}💾 Token cacheado com segurança${NC}" >&2
            
            echo "$new_token"
            return 0
        else
            echo -e "${RED}❌ Token inválido${NC}" >&2
            return 1
        fi
    fi
    
    return 1
}

# Função para limpar cache
clear_cache() {
    rm -f "${TOKEN_CACHE_ENCRYPTED}"
    echo -e "${GREEN}✅ Cache de token limpo${NC}"
}

# Função para mostrar status
show_status() {
    echo -e "${BLUE}🔑 Status de Autenticação GitHub${NC}"
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    
    # Verificar variável de ambiente
    if [ -n "$GITHUB_TOKEN" ]; then
        if validate_token "$GITHUB_TOKEN"; then
            local username=$(get_user_info "$GITHUB_TOKEN")
            echo -e "${GREEN}✅ GITHUB_TOKEN: Válido (${username})${NC}"
        else
            echo -e "${RED}❌ GITHUB_TOKEN: Inválido${NC}"
        fi
    else
        echo -e "${YELLOW}⚠️  GITHUB_TOKEN: Não definido${NC}"
    fi
    
    # Verificar cache
    local cached_token=$(decrypt_token 2>/dev/null)
    if [ -n "$cached_token" ]; then
        if validate_token "$cached_token"; then
            local username=$(get_user_info "$cached_token")
            echo -e "${GREEN}✅ Cache: Válido (${username})${NC}"
        else
            echo -e "${RED}❌ Cache: Expirado${NC}"
        fi
    else
        echo -e "${YELLOW}⚠️  Cache: Vazio${NC}"
    fi
    
    # Verificar gh CLI
    if command -v gh &> /dev/null && gh auth status &> /dev/null; then
        local gh_token=$(gh auth token 2>/dev/null)
        if [ -n "$gh_token" ] && validate_token "$gh_token"; then
            local username=$(get_user_info "$gh_token")
            echo -e "${GREEN}✅ gh CLI: Autenticado (${username})${NC}"
        else
            echo -e "${RED}❌ gh CLI: Token inválido${NC}"
        fi
    else
        echo -e "${YELLOW}⚠️  gh CLI: Não autenticado${NC}"
    fi
    
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
}

# Comando principal
case "${1:-get}" in
    get)
        get_github_token
        ;;
    clear)
        clear_cache
        ;;
    status)
        show_status
        ;;
    *)
        echo "Uso: $0 {get|clear|status}"
        exit 1
        ;;
esac
