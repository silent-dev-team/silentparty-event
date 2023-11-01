package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/silent-dev-team/silentparty-event/pocketbase/utils"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		settings, _ := dao.FindSettings()
		settings.Meta.AppName = "Silent App"
		settings.Meta.AppUrl = "https://app.silentparty-hannver.de"
		settings.Logs.MaxDays = 30

		settings.Meta.SenderName = utils.Getenv("SMTP_SENDER_NAME")    //"Silent Party Hannover"
		settings.Meta.SenderAddress = utils.Getenv("SMTP_SENDER_MAIL") //"info@silentparty-hannover.de"
		settings.Smtp.Enabled = false                                  // true
		settings.Smtp.Host = utils.Getenv("SMTP_HOST")                 //"smtp.goneo.de"
		settings.Smtp.Port = 587
		settings.Smtp.Username = utils.Getenv("SMTP_USERNAME") //"info@silentparty-hannover.de"
		settings.Smtp.Password = utils.Getenv("SMTP_PW")

		settings.Backups.Cron = "0 */6 * * *"
		settings.Backups.CronMaxKeep = 30

		settings.Backups.S3.Enabled = true
		settings.Backups.S3.Endpoint = "http://s3.eu-central-1.amazonaws.com"
		settings.Backups.S3.Bucket = "silentpartybackup"
		settings.Backups.S3.Region = "eu-central-1"
		settings.Backups.S3.AccessKey = utils.Getenv("S3_ACCESS_KEY")
		settings.Backups.S3.Secret = utils.Getenv("S3_SECRET")
		settings.Backups.S3.ForcePathStyle = false

		return dao.SaveSettings(settings)
	}, nil)
}
