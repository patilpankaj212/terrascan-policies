provider "azurerm" {
  features {}
}

data "azurerm_client_config" "current" {
}

resource "azurerm_resource_group" "resource_group_details" {
  name     = "my-resource-group"
  location = "West US"
}

resource "random_id" "server" {
  keepers = {
    ami_id = 1
  }

  byte_length = 8
}

resource "azurerm_key_vault" "vault_details" {
  name                = "keyvaultkeyexample"
  location            = azurerm_resource_group.resource_group_details.location
  resource_group_name = azurerm_resource_group.resource_group_details.name
  tenant_id           = data.azurerm_client_config.current.tenant_id

  sku_name = "premium"

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    key_permissions = [
      "create",
      "get",
    ]

    secret_permissions = [
      "set",
    ]
  }

  tags = {
    environment = "Production"
  }
}

resource "azurerm_key_vault_key" "ckeckKeyExpirationIsSet" {
  name         = "generated-certificate"
  key_vault_id = azurerm_key_vault.vault_details.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
  ]
}
