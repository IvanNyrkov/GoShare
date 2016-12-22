package main

import (
	"github.com/IvanNyrkov/GoShare/src/api"
	"github.com/IvanNyrkov/GoShare/src/rand/sentence"
)

// InitModules injects all required app modules
func (app *App) InitModules() error {
	var err error
	if app.RandSentenceModule, err = sentence.NewModule(sentence.DefaultConfig); err != nil {
		return err
	}
	app.APIModule = api.NewModule(api.ModuleConfig{})
	return nil
}
