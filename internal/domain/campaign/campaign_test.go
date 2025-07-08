package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert" // Usando para melhorar a legibilidade do código, retirando ifs
)

// Constante sendo utilizada no código
var (
	name     = "Campaign x"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

/*
	Para fazer um teste de unidade você precisa ter noção sobre o que é os 3 As:
	Arrange: Preparar o ambiente para o teste ( Organizar o teste)
	Act: Executar a ação que está sendo testada
	Assert: Verificar se o resultado é o esperado
*/

func TestNewCampaign_CreateCampaign(t *testing.T) {

	// Arrange
	assert := assert.New(t)

	// Act
	campaign := NewCampaign(name, content, contacts)

	// Assert
	println(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaing_IDIsNotNill(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign := NewCampaign(name, content, contacts)

	// Assert
	assert.NotNil(campaign.ID)
}

func Test_NewCampaing_CreateOnMustBeNow(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	// Act
	campaign := NewCampaign(name, content, contacts)

	// Assert
	assert.Greater(campaign.CreatedOn, now) // Lembrar que valores de data sempre vão retornar um valor, ou seja, não valide para que caso ele seja nulo
}
