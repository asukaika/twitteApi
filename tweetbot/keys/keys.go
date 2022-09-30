package keys

import (
	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi() *anaconda.TwitterApi {
    anaconda.SetConsumerKey("VBrgNdckhllQGHSu3WRbxbKKs")
    anaconda.SetConsumerSecret("6NWSu5EHLvAuNfNwx3BrMZTMWzYrJNAoQKXmuhd13OkRKLgCX0")
    api := anaconda.NewTwitterApi("1575768596911120384-kUCnjm42wZbH2dSDbZgBpPomr9OurF", "BJyC5pbsYmzvCN7RLdcNZ0g3t38PelDGpKP4mb4PLkrwr")
    return api
}