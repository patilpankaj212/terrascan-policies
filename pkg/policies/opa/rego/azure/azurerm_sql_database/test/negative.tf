provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "a_rg" {
  name     = "acceptanceTestResourceGroup1"
  location = "West US"
}

resource "azurerm_sql_server" "a_sql" {
  name                         = "mysqlserver"
  resource_group_name          = azurerm_resource_group.a_rg.name
  location                     = "West US"
  version                      = "12.0"
  administrator_login          = "4dm1n157r470r"
  administrator_login_password = "4-v3ry-53cr37-p455w0rd"
}

resource "azurerm_sql_database" "checkAuditEnabled" {
  name                = "mysqldatabase"
  resource_group_name = azurerm_resource_group.a_rg.name
  location            = "West US"
  server_name         = azurerm_sql_server.a_sql.name

  tags = {
    environment = "production"
  }

  threat_detection_policy {
    state = "Enabled"
  }
}
