package xlsxDecoder

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {

}

func TestGetURLForDownloading(t *testing.T) {
	url, err := GetURLForDownloading("avito-test/table.xlsx", "OAuth AgAAAAA1z4O5AADLW7ibSa25TUIVocRFVAYdP1Q")
	assert.NoError(t, err)
	assert.NotEqual(t, url, "")
}

func TestDownloadFile(t *testing.T) {
	url := "https://downloader.disk.yandex.ru/disk/09d764dc719df903049c1b6f68a7f2b2e58fd4b02b622f10cc5d9f7e7801c38e/600ad4eb/t_4Hiwt2Y9N1kryGs2lhE5eJzsDt5uNcNgBt6qGyl0xg9oI7Kuf1IWmljWXeCCAlIxW7bcPeO_nxp1OAJ_LDFA%3D%3D?uid=902792121&filename=table.xlsx&disposition=attachment&hash=&limit=0&content_type=application%2Fvnd.openxmlformats-officedocument.spreadsheetml.sheet&owner_uid=902792121&fsize=8542&hid=69d767ba52a77805da92ef3433e0ff6b&media_type=document&tknv=v2&etag=6e51693ba0d536719de929634dfde2fd"
	err := DownloadFile(url, "testingFile")
	assert.NoError(t, err)
	assert.FileExists(t, "xlsxFiles/testingFile.xlsx")
	os.Remove("xlsxFiles/testingFile.xlsx")
}