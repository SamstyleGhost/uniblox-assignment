package tests

import (
	"fmt"
	"testing"

	jsonfileworker "github.com/SamstyleGhost/uniblox-assignment/lib/json-file-worker"
	"github.com/SamstyleGhost/uniblox-assignment/models"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func TestGetAllObjects(t *testing.T) {
	var users []models.User

	err := jsonfileworker.GetAllObjects("../data/users.json", &users)
	if err != nil {
		t.Error(err)
	}

	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		t.Error(err)
	}
	fmt.Print(string(b))
}

func TestSetAllObjects(t *testing.T) {
	var coupons []models.Coupon

	err := jsonfileworker.GetAllObjects("../data/coupons.json", &coupons)
	if err != nil {
		fmt.Println("Errror in GetAllObjects")
		t.Error(err)
	}

	newCoupon := new(models.Coupon)
	newCoupon.UserID, _ = uuid.Parse("d9cd8e94-c68c-4fcc-9a83-cd51f42ffcbe")
	newCoupon.CouponCode, _ = uuid.Parse("d9cd8e94-c68c-4fcc-9a83-cd51f42ffcbe")
	coupons = append(coupons, *newCoupon)

	err = jsonfileworker.SetAllObjects("../data/coupons.json", &coupons)
	if err != nil {
		t.Error(err)
	}

	err = jsonfileworker.GetAllObjects("../data/coupons.json", &coupons)
	if err != nil {
		fmt.Println("Errror in GetAllObjects")
		t.Error(err)

		b, err := json.MarshalIndent(coupons, "", "  ")
		if err != nil {
			t.Error(err)
		}
		fmt.Print(string(b))
	}
}
