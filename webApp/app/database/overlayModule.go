package database

import (
	"context"
)

type OverlayModuleType int

const (
	InvalidModule = OverlayModuleType(iota)
	AlertBox
)

type OverlayModule struct {
	ID     int
	Type   OverlayModuleType
	Top    int
	Left   int
	Width  int
	Height int
	IsNew  bool
}

func (overlay *OverlayStruct) GetModules() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	rows, err := conn.Query(context.Background(), `SELECT id, "type", "top", "left", "width", "height" FROM public."overlayModules" WHERE "overlayID"=$1`, overlay.ID)
	if err != nil {
		return err
	}

	if overlay.ModuleInfo == nil {
		overlay.ModuleInfo = make(map[int]*OverlayModule)
	}

	for rows.Next() {
		module := &OverlayModule{}
		err = rows.Scan(&module.ID, &module.Type, &module.Top, &module.Left, &module.Width, &module.Height)
		if err != nil {
			return err
		}
		overlay.ModuleInfo[overlay.ID] = module
	}

	return nil
}

func (overlay *OverlayStruct) NewModule(Type OverlayModuleType, top, left, width, height int) (int, error) {
	conn, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	module := &OverlayModule{
		Type:   Type,
		Top:    top,
		Left:   left,
		Width:  width,
		Height: height,
	}

	err = conn.QueryRow(context.Background(),
		`INSERT INTO public."overlayModules"("overlayID", "type", "top", "left", "width", "height") VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id";`,
		overlay.ID, module.Type, module.Top, module.Left, module.Width, module.Height).Scan(&module.ID)

	if err != nil {
		return module.ID, err
	}

	if overlay.ModuleInfo == nil {
		overlay.ModuleInfo = make(map[int]*OverlayModule)
	}

	overlay.ModuleInfo[overlay.ID] = module
	return module.ID, nil
}

func (module *OverlayModule) Save() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(context.Background(), `UPDATE public."overlayModules" SET "top"=$1, "left"=$2, "width"=$3, "height"=$4 WHERE "id"=$5`,
		module.Top, module.Left, module.Width, module.Height, module.ID)
	return err
}

func (module *OverlayModule) Delete() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(context.Background(), `DELETE FROM public."overlayModules" WHERE "id"=$1`, module.ID)
	if err != nil {
		return err
	}

	module = nil
	return nil
}

func (module *OverlayModule) Update(top, left, width, height int) error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	module.Top = top
	module.Left = left
	module.Width = width
	module.Height = height

	_, err = conn.Exec(context.Background(), `UPDATE public."overlayModules" SET "top"=$1, "left"=$2, "width"=$3, "height"=$4 WHERE "id"=$5`,
		module.Top, module.Left, module.Width, module.Height, module.ID)
	return err
}
