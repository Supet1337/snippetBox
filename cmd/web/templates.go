package main

import (
	"github.com/Supet1337/snippetBox/pkg/models"
	"time"
)

// Создаем тип templateData, который будет действовать как хранилище для
// любых динамических данных, которые нужно передать HTML-шаблонам.

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
	Time     time.Time
}
