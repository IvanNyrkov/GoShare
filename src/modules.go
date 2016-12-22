package main

import (
	"github.com/IvanNyrkov/GoShare/src/api"
	"github.com/IvanNyrkov/GoShare/src/rand/sentence"
	"github.com/IvanNyrkov/GoShare/src/store"
)

// InitModules injects all required app modules
func (app *App) InitModules() error {
	// Random Sentence Module
	var err error
	if app.RandSentenceModule, err = sentence.NewModule(sentence.DefaultConfig); err != nil {
		return err
	}
	// File Storage Module
	app.FileStorageModule = store.NewModule(store.DefaultConfig)
	// API Module
	app.APIModule = api.NewModule(api.ModuleConfig{
		RandSentenceService: app.RandSentenceModule.GetService(),
		FileStorageService:  app.FileStorageModule.GetService(),
	})
	return nil
}
