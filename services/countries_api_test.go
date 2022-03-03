package services

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"world-api/entities"
	"world-api/utils"
)

func TestCountriesAPIImpl_GetCountry(t *testing.T) {
	type fields struct {
		MakeRequest utils.MakeAPIRequest
	}
	type args struct {
		countryName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Country
		wantErr bool
	}{
		{
			name: "server error",
			fields: fields{
				MakeRequest: func(context.Context, string) ([]byte, error) {
					return nil, errors.New("server error")
				},
			},
			args: args{
				countryName: "111",
			},
			wantErr: true,
		},

		{
			name: "correct country name",
			fields: fields{MakeRequest: func(context.Context, string) ([]byte, error) {
				return []byte("[{\"name\":{\"common\":\"Brazil\",\"official\":\"Federative Republic of Brazil\",\"nativeName\":{\"por\":{\"official\":\"República Federativa do Brasil\",\"common\":\"Brasil\"}}},\"tld\":[\".br\"],\"cca2\":\"BR\",\"ccn3\":\"076\",\"cca3\":\"BRA\",\"cioc\":\"BRA\",\"independent\":true,\"status\":\"officially-assigned\",\"unMember\":true,\"currencies\":{\"BRL\":{\"name\":\"Brazilian real\",\"symbol\":\"R$\"}},\"idd\":{\"root\":\"+5\",\"suffixes\":[\"5\"]},\"capital\":[\"Brasília\"],\"altSpellings\":[\"BR\",\"Brasil\",\"Federative Republic of Brazil\",\"República Federativa do Brasil\"],\"region\":\"Americas\",\"subregion\":\"South America\",\"languages\":{\"por\":\"Portuguese\"},\"translations\":{\"ara\":{\"official\":\"جمهورية البرازيل الاتحادية\",\"common\":\"البرازيل\"},\"ces\":{\"official\":\"Brazilská federativní republika\",\"common\":\"Brazílie\"},\"cym\":{\"official\":\"Gweriniaeth Ffederal Brasil\",\"common\":\"Brasil\"},\"deu\":{\"official\":\"Föderative Republik Brasilien\",\"common\":\"Brasilien\"},\"est\":{\"official\":\"Brasiilia Liitvabariik\",\"common\":\"Brasiilia\"},\"fin\":{\"official\":\"Brasilian liittotasavalta\",\"common\":\"Brasilia\"},\"fra\":{\"official\":\"République fédérative du Brésil\",\"common\":\"Brésil\"},\"hrv\":{\"official\":\"Savezne Republike Brazil\",\"common\":\"Brazil\"},\"hun\":{\"official\":\"Brazil Szövetségi Köztársaság\",\"common\":\"Brazília\"},\"ita\":{\"official\":\"Repubblica federativa del Brasile\",\"common\":\"Brasile\"},\"jpn\":{\"official\":\"ブラジル連邦共和国\",\"common\":\"ブラジル\"},\"kor\":{\"official\":\"브라질 연방 공화국\",\"common\":\"브라질\"},\"nld\":{\"official\":\"Federale Republiek Brazilië\",\"common\":\"Brazilië\"},\"per\":{\"official\":\"جمهوری فدراتیو برزیل\",\"common\":\"برزیل\"},\"pol\":{\"official\":\"Federacyjna Republika Brazylii\",\"common\":\"Brazylia\"},\"por\":{\"official\":\"República Federativa do Brasil\",\"common\":\"Brasil\"},\"rus\":{\"official\":\"Федеративная Республика Бразилия\",\"common\":\"Бразилия\"},\"slk\":{\"official\":\"Brazílska federatívna republika\",\"common\":\"Brazília\"},\"spa\":{\"official\":\"República Federativa del Brasil\",\"common\":\"Brasil\"},\"swe\":{\"official\":\"Förbundsrepubliken Brasilien\",\"common\":\"Brasilien\"},\"urd\":{\"official\":\"وفاقی جمہوریہ برازیل\",\"common\":\"برازیل\"},\"zho\":{\"official\":\"巴西联邦共和国\",\"common\":\"巴西\"}},\"latlng\":[-10.0,-55.0],\"landlocked\":false,\"borders\":[\"ARG\",\"BOL\",\"COL\",\"GUF\",\"GUY\",\"PRY\",\"PER\",\"SUR\",\"URY\",\"VEN\"],\"area\":8515767.0,\"demonyms\":{\"eng\":{\"f\":\"Brazilian\",\"m\":\"Brazilian\"},\"fra\":{\"f\":\"Brésilienne\",\"m\":\"Brésilien\"}},\"flag\":\"\\uD83C\\uDDE7\\uD83C\\uDDF7\",\"maps\":{\"googleMaps\":\"https://goo.gl/maps/waCKk21HeeqFzkNC9\",\"openStreetMaps\":\"https://www.openstreetmap.org/relation/59470\"},\"population\":212559409,\"gini\":{\"2019\":53.4},\"fifa\":\"BRA\",\"car\":{\"signs\":[\"BR\"],\"side\":\"right\"},\"timezones\":[\"UTC-05:00\",\"UTC-04:00\",\"UTC-03:00\",\"UTC-02:00\"],\"continents\":[\"South America\"],\"flags\":{\"png\":\"https://flagcdn.com/w320/br.png\",\"svg\":\"https://flagcdn.com/br.svg\"},\"coatOfArms\":{\"png\":\"https://mainfacts.com/media/images/coats_of_arms/br.png\",\"svg\":\"https://mainfacts.com/media/images/coats_of_arms/br.svg\"},\"startOfWeek\":\"monday\",\"capitalInfo\":{\"latlng\":[-15.79,-47.88]},\"postalCode\":{\"format\":\"#####-###\",\"regex\":\"^(\\\\d{8})$\"}}]"), nil
			}},
			args: args{countryName: "Brazil"},
			want: entities.Country{
				Name: entities.Name{
					Common:   "Brazil",
					Official: "Federative Republic of Brazil",
					NativeName: entities.NativeName{
						Por: entities.Por{
							Official: "República Federativa do Brasil",
							Common:   "Brasil",
						},
					},
				},
				Tld:         []string{".br"},
				Cca2:        "BR",
				Ccn3:        "076",
				Cca3:        "BRA",
				Cioc:        "BRA",
				Independent: true,
				Status:      "officially-assigned",
				UnMember:    true,
				Currencies: entities.Currencies{
					Brl: entities.Brl{
						Name:   "Brazilian real",
						Symbol: "R$",
					},
				},
				Idd: entities.Idd{
					Root:     "+5",
					Suffixes: []string{"5"},
				},
				Capital: []string{"Brasília"},
				AltSpellings: []string{
					"BR",
					"Brasil",
					"Federative Republic of Brazil",
					"República Federativa do Brasil",
				},
				Region:    "Americas",
				Subregion: "South America",
				Languages: entities.Languages{
					Por: "Portuguese",
				},
				Translations: entities.Translations{
					Ara: entities.TranslationsItem{
						Official: "جمهورية البرازيل الاتحادية",
						Common:   "البرازيل",
					},
					Ces: entities.TranslationsItem{
						Official: "Brazilská federativní republika",
						Common:   "Brazílie",
					},
					Cym: entities.TranslationsItem{
						Official: "Gweriniaeth Ffederal Brasil",
						Common:   "Brasil",
					},
					Deu: entities.TranslationsItem{
						Official: "Föderative Republik Brasilien",
						Common:   "Brasilien",
					},
					Est: entities.TranslationsItem{
						Official: "Brasiilia Liitvabariik",
						Common:   "Brasiilia",
					},
					Fin: entities.TranslationsItem{
						Official: "Brasilian liittotasavalta",
						Common:   "Brasilia",
					},
					Fra: entities.TranslationsItem{
						Official: "République fédérative du Brésil",
						Common:   "Brésil",
					},
					Hrv: entities.TranslationsItem{
						Official: "Savezne Republike Brazil",
						Common:   "Brazil",
					},
					Hun: entities.TranslationsItem{
						Official: "Brazil Szövetségi Köztársaság",
						Common:   "Brazília",
					},
					Ita: entities.TranslationsItem{
						Official: "Repubblica federativa del Brasile",
						Common:   "Brasile",
					},
					Jpn: entities.TranslationsItem{
						Official: "ブラジル連邦共和国",
						Common:   "ブラジル",
					},
					Kor: entities.TranslationsItem{
						Official: "브라질 연방 공화국",
						Common:   "브라질",
					},
					Nld: entities.TranslationsItem{
						Official: "Federale Republiek Brazilië",
						Common:   "Brazilië",
					},
					Per: entities.TranslationsItem{
						Official: "جمهوری فدراتیو برزیل",
						Common:   "برزیل",
					},
					Pol: entities.TranslationsItem{
						Official: "Federacyjna Republika Brazylii",
						Common:   "Brazylia",
					},
					Por: entities.TranslationsItem{
						Official: "República Federativa do Brasil",
						Common:   "Brasil",
					},
					Rus: entities.TranslationsItem{
						Official: "Федеративная Республика Бразилия",
						Common:   "Бразилия",
					},
					Slk: entities.TranslationsItem{
						Official: "Brazílska federatívna republika",
						Common:   "Brazília",
					},
					Spa: entities.TranslationsItem{
						Official: "República Federativa del Brasil",
						Common:   "Brasil",
					},
					Swe: entities.TranslationsItem{
						Official: "Förbundsrepubliken Brasilien",
						Common:   "Brasilien",
					},
					Urd: entities.TranslationsItem{
						Official: "وفاقی جمہوریہ برازیل",
						Common:   "برازیل",
					},
					Zho: entities.TranslationsItem{
						Official: "巴西联邦共和国",
						Common:   "巴西",
					},
				},
				Latlng: []float32{
					-10,
					-55,
				},
				Landlocked: false,
				Borders: []string{
					"ARG",
					"BOL",
					"COL",
					"GUF",
					"GUY",
					"PRY",
					"PER",
					"SUR",
					"URY",
					"VEN",
				},
				Area: 8515767,
				Demonyms: entities.Demonyms{
					Eng: entities.DemonymsItem{
						F: "Brazilian",
						M: "Brazilian",
					},
					Fra: entities.DemonymsItem{
						F: "Brésilienne",
						M: "Brésilien",
					},
				},
				Flag: "🇧🇷",
				Maps: entities.Maps{
					GoogleMaps:     "https://goo.gl/maps/waCKk21HeeqFzkNC9",
					OpenStreetMaps: "https://www.openstreetmap.org/relation/59470",
				},
				Population: 212559409,
				Gini: entities.Gini{
					Num2019: 53.4,
				},
				Fifa: "BRA",
				Car: entities.Car{
					Signs: []string{"BR"},
					Side:  "right",
				},
				Timezones: []string{
					"UTC-05:00",
					"UTC-04:00",
					"UTC-03:00",
					"UTC-02:00",
				},
				Continents: []string{"South America"},
				Flags: entities.Flags{
					Png: "https://flagcdn.com/w320/br.png",
					Svg: "https://flagcdn.com/br.svg",
				},
				CoatOfArms: entities.CoatOfArms{
					Png: "https://mainfacts.com/media/images/coats_of_arms/br.png",
					Svg: "https://mainfacts.com/media/images/coats_of_arms/br.svg",
				},
				StartOfWeek: "monday",
				CapitalInfo: entities.CapitalInfo{
					Latlng: []float64{
						-15.79,
						-47.88,
					},
				},
				PostalCode: entities.PostalCode{
					Format: "#####-###",
					Regex:  "^(\\d{8})$",
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CountriesAPIImpl{
				MakeRequest: tt.fields.MakeRequest,
			}
			ctx := context.TODO()
			got, err := c.GetCountry(ctx, tt.args.countryName)

			if (err != nil) != tt.wantErr {
				t.Errorf("CountriesAPIImpl.GetCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountriesAPIImpl.GetCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
