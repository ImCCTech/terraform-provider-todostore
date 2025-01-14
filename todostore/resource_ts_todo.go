package todostore

import (
	"strconv"

	sdk "github.com/ImCCTech/go-todostore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTSTodo() *schema.Resource {
	return &schema.Resource{
		Create: resourceTSTodoCreate,
		Read:   resourceTSTodoRead,
		Update: resourceTSTodoUpdate,
		Delete: resourceTSTodoDelete,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"memo": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceTSTodoCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*sdk.Client)
	options := sdk.TodoCreateOptions{
		Title: d.Get("title").(string),
		Memo:  d.Get("memo").(string),
	}

	todo, err := conn.Todos.Create(options)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(todo.ID))
	resourceTSTodoRead(d, meta)
	return nil
}

func resourceTSTodoRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*sdk.Client)
	todo_id, _ := strconv.Atoi(d.Id())
	todo, err := conn.Todos.Read(todo_id)
	if err != nil {
		return err
	}
	d.Set("title", todo.Title)
	d.Set("memo", todo.Memo)
	return nil
}

func resourceTSTodoUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*sdk.Client)
	options := sdk.TodoUpdateOptions{}
	if d.HasChange("title") {
		options.Title = d.Get("title").(string)
	}
	if d.HasChange("memo") {
		options.Memo = d.Get("memo").(string)
	}
	todo_id, _ := strconv.Atoi(d.Id())
	conn.Todos.Update(todo_id, options)
	return resourceTSTodoRead(d, meta)
}

func resourceTSTodoDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*sdk.Client)
	todo_id, _ := strconv.Atoi(d.Id())
	conn.Todos.Delete(todo_id)
	return nil
}
