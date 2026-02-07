package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"x/llm"
)

func recommendations_handler(w http.ResponseWriter, r *http.Request) {
	// shouldMock := r.URL.Query().Get("mock")
	// if shouldMock == "1" {
		time.Sleep(time.Second * 4)
		mockResult := `{"data":[{"name":"Thermalright ASF Black V2 AM5 CPU Holder, AM5 Safety Fixed Frame, Corrective Anti-Bending Fixing Frame","image_path":"https://m.media-amazon.com/images/I/51TDFYjTZ7L._AC_UY218_.jpg","provider_url":"https://www.amazon.com/Thermalright-ASF-V2-Corrective-Anti-Bending/dp/B0CYQ3LDML/ref=sr_1_2?dib=eyJ2IjoiMSJ9.JzFPf3kV6_qaGxYeA0R0euIPKvMvBTkzuqvDVuozqI8tJZwk-ufngysZGQYv22sXyU8wdlAaZuMLyaGILQSMc97Xe6NdVeiT-6j_C0gwzVE7I7dNcEVRP_SWzQeuT-JLmiYdRRGu4z9vahnITAssbjS9nw7zzXgTMTSgtDqi7ZgZuA5wreBErmAClVi00ZEncIEqim4-TvhA5WvzLLcApxN4eRphE-qM-thRhP_SgsQ.aZQFmYedj6igeTH-1F4BqCtJleZhb_eQX1rnGc20Ias\u0026dib_tag=se\u0026keywords=AM5%2Bcontact%2Bframe%2Banti-bending%2Bbracket%2Bfor%2BRyzen%2B7000%2BAM5\u0026qid=1770505081\u0026sr=8-2","price":"$8.69","error":""},{"name":"be quiet! MC1 Pro M.2 SSD Cooler, Heatpipe Heatsink for 2280 NVMe, Dual-Sided Compatible","image_path":"https://m.media-amazon.com/images/I/61mBgqI2StL._AC_UY218_.jpg","provider_url":"https://www.amazon.com/sspa/click?ie=UTF8\u0026spc=MTo1NjQzMDY3MjIwNjU4MDMwOjE3NzA1MDUwODU6c3BfYXRmOjMwMDk4ODcwMzUwNDYwMjo6MDo6\u0026url=%2Fquiet-BZ003-Cooler-heatsink-modules%2Fdp%2FB08YRVM51Q%2Fref%3Dsr_1_1_sspa%3Fdib%3DeyJ2IjoiMSJ9.dRZlSJm_8RHRU1ora6L9uIGdynWE27qubBHUoIWGZjE_vpooPqYv1iY7cbJi_KXXZTh53u2MC84kG4YMovPAAa885CVsxXJEyqQbfZsSIInvQ1zsIUTDYoyJ0ayL_XlhnpwgtA5L8cmk4d2StypA8YmoatDJ8Pxo4N_VwCr4aKIZKbfa5O0nnhW-8o5iemqPsBfs3zgZqo53j4hnzWfJ92jHnMDiCZQCCA9hqLVnAkg.qBSAmehW69gFz4Z_MNDjMAgvGKBj0cVpadEAjkzVFfQ%26dib_tag%3Dse%26keywords%3DNVMe%252BM.2%252Bheatsink%252Bwith%252Bthermal%252Bpads%252BPCIe%252B4.0%252B2280%26qid%3D1770505085%26sr%3D8-1-spons%26sp_csd%3Dd2lkZ2V0TmFtZT1zcF9hdGY%26psc%3D1","price":"$19.90","error":""},{"name":"CableMod Universal Pro ModMesh Sleeved 12V-2x6 12VHPWR StealthSense Direct PCIe Cable (Black, 16-pin to 16-pin, 60cm)","image_path":"https://m.media-amazon.com/images/I/51g8i+F5ciL._AC_UY218_.jpg","provider_url":"https://www.amazon.com/CableMod-Universal-ModMesh-Sleeved-StealthSense/dp/B0DSHQRG2T/ref=sr_1_3?dib=eyJ2IjoiMSJ9.NzGNttnTyG2en0CwLHoNQSQCau8rQQYHix7Ob8F_4v3Qzp-d4wrL7043cLE_eTPTXbm6sxg5F2y083cLMUrN4OeAY9Va4U0n2-XqlqQpN9FMaOxt_eQLu7Hry8xydLq6eruTF0PzjqDFcyujr7PToHw3AfKVxx3RGfPzi-7IsYWXn59nNKcbJfE0Sb4bPnRXpcQJN6ZyzJZLrodC_TT1UJH-h3aiwfCqJmJ2DYU5bVU.jpzd9M3zVSecu21OFKVgbwGIcpE-Lc9fePjL3JSRVow\u0026dib_tag=se\u0026keywords=ATX%2B3.1%2B12V-2x6%2Bcable%2Bcombs%2Bsleeved%2Bextension%2Bkit\u0026qid=1770505083\u0026sr=8-3","price":"$24.90","error":""},{"name":"2 Pack 51mm Espresso Puck Screen, Reusable 1.7mm Thickness 150μm Coffee Filter Mesh Plate","image_path":"https://m.media-amazon.com/images/I/81LY0QCxmOL._AC_UL320_.jpg","provider_url":"https://www.amazon.com/Espresso-Screen-Reusable-Thickness-Portafilter/dp/B09XLMK5LR/ref=sr_1_1?dib=eyJ2IjoiMSJ9.BtgA92MxV1qplTHA96v6uJkdWNwyKA_LWzGcB1yxmKJjgj8wZ0KsrLCusi1E2aLtug5qAgWvY5DyruTCEaNnsvtqWXayh-T4Hy-WYw-B2oKikGQXsxDf6RWIYe27Sdp6jXUEshe8kMGXTgJ3S2yxntF2L1ikB3DHHGb91jmQE1wJ__PambBOPKulf4SJhzacxJ1N87woibjmwTjVplvhsfijRg2vvbUKIE_6wTegUSB1rByqThC7tjICUltdhRRMQ1hLYrXGnL81cR3ptmDzihCIsgxYKrdGfmr4Rlf8Bj0.Wc_OPzMgJZlvrNjfA8g0CHDfjC32czUdkqaWhmUCEbQ\u0026dib_tag=se\u0026keywords=51mm%2Bespresso%2Bpuck%2Bscreen%2B1.7mm%2Bstainless%2Bsteel\u0026qid=1770505086\u0026sr=8-1","price":"$6.99","error":""},{"name":"Kasa Smart Plug Power Strip HS300, Surge Protector with 6 Individually Controlled Smart Outlets and 3 USB Ports","image_path":"https://m.media-amazon.com/images/I/61PI8akrKOL._AC_UY218_.jpg","provider_url":"https://www.amazon.com/Kasa-Smart-Power-Strip-TP-Link/dp/B07G95FFN3/ref=sr_1_1?dib=eyJ2IjoiMSJ9.FykPj9PLN6f3awtsf9hbkdQuQoodtpLIZxl2yWOVqO2jbMeAADFElVNFfBUcv2WhEvnIWMFZaRDVTq2YuiFC2vF5RBcG2lXDd57qjXhOyhs7wnuscM_mSutxfSn6-sKcOwtiRP2RHgCTj5R5rZHNHhHfZqti2uEYdzng13xqpSosQRd-1PQLjB1YP97wgq3TgruOy6U4oCz8aWrFyJ7NocSnBIXKSQf-kShy31XBWg8.1C2GrMrcPNO5AZ0_iUlQnhciZr85yCs0fK1axC4oazA\u0026dib_tag=se\u0026keywords=smart%2Bpower%2Bstrip%2Bsurge%2Bprotector%2Benergy%2Bmonitoring%2BAlexa\u0026qid=1770505089\u0026sr=8-1","price":"$39.98","error":""},{"name":"VIVO Under Desk 17 inch Cable Management Tray, Power Strip Holder, Cord Organizer (DESK-AC06-1C)","image_path":"https://m.media-amazon.com/images/I/71CoiO1qrtL._AC_UY218_.jpg","provider_url":"https://www.amazon.com/VIVO-Management-Holder-Organizer-DESK-AC06-1C/dp/B089B4XZM4/ref=sr_1_5?dib=eyJ2IjoiMSJ9.gLOp4wYe18t5CCQld9UiJ-1uHDJTcFxniDFEfddVj0usgUzCY5HNpqQk1_0HAJ5qDmHmz8VPxtuXP4a2XHV2bPNXEEizDCACMG5hH6LIfhWhdCIzO2ywZXq0gjWC1VCk2bObc5GKM5ZPx9qYrJgpFWrEOg6LABAR9xk6SpNQD1eoQ-RFmodESiXWpmGh18tw3JpkyK1_ROsGjb8-Wrvmm2DO3Qyf0NsTas9p3LKA9Jk.HqU2OV7AWEBav-Y97tkFVucWKwXVMnjnwQkPGi_mRuk\u0026dib_tag=se\u0026keywords=desk%2Bcable%2Bmanagement%2Btray%2Bunder%2Bdesk%2Bmetal\u0026qid=1770505090\u0026sr=8-5","price":"$19.99","error":""},{"name":"Pulsar Gaming Gears - Micro Bungee ES: Drag-Free Wired Mouse Support","image_path":"https://m.media-amazon.com/images/I/618hPVlV5NL._AC_UY218_.jpg","provider_url":"https://www.amazon.com/Pulsar-Gaming-Gears-Drag-Free-Support/dp/B0C81FSVCC/ref=sr_1_1?dib=eyJ2IjoiMSJ9.PRjXphJbmWaY05lfvjjyim0tRRjzW9qrwaJfnreiIe1Lb7hfIOaGwEex7KDMrOa2nV6z12IVRbkEJejI0KedDeZY2sO1bIHHpfbzDtv8IxkWltz4RgMyX0hqj1BA_anOBM2XCSnIhsDdf5v62jd1HumXY4XNogRFcXL9_fKWUu674NI0DFsPk6L1mBnX86Zma1ltDJMdJm5lsFXGwUsJKRaIotudrIg4HMAc9C3DIQM.8Jmh7Q9pJ4tq2chY7gTpEO4OnRUh0y9w6ymTOFa4ZgA\u0026dib_tag=se\u0026keywords=gaming%2Bmouse%2Bbungee%2Bfor%2Bwired%2Bmouse\u0026qid=1770505095\u0026sr=8-1","price":"$6.95","error":""},{"name":"Advantage360 Palm Pads - Magnetic | Cushioned Foam | Washable","image_path":"https://m.media-amazon.com/images/I/61xLad0wlEL._AC_UY218_.jpg","provider_url":"https://www.amazon.com/Advantage360-Palm-Pads-Magnetic-Cushioned/dp/B0BCHFSRN2/ref=sr_1_1_mod_primary_new?dib=eyJ2IjoiMSJ9.JVRWejCNFBCqMXpaE6XCroxSYYwODjBnW1f1qcCZ0sb0LcZFsgds1pnTujOw-WcXY-XgyeTuS8QrSihX5Allvj9XDkwFMljhVIrb-PDk7X92xvrxF9jg-MZOo68ansGApzXwLV6wFOR9GU6S7GilJRAMLx59LTELjpav3je1Z6ij8piRa1X_4WIQLoU_1DSJWHrzK54p84F3SHLVtD_Byrqd6Dl2Os-C0G7M84eaCvw.0xMs7KZf05LYevacmjc1jYdFTf3ilrfDZRsOWk0llS0\u0026dib_tag=se\u0026keywords=Kinesis%2BAdvantage360%2Bpalm%2Bpads%2Breplacement\u0026qid=1770505096\u0026sbo=RZvfv%2F%2FHxDF%2BO5021pAnSA%3D%3D\u0026sr=8-1","price":"$25.99","error":""},{"name":"Go Hang It! Pro, All-in-One Picture Hanging Kit - Picture Leveling and Hanging Tool (85-Piece)","image_path":"https://m.media-amazon.com/images/I/71tSEvkKyLL._AC_UL320_.jpg","provider_url":"https://www.amazon.com/Go-Hang-Pro-All-One/dp/B07TQKT16K/ref=sr_1_1?dib=eyJ2IjoiMSJ9.23pBFjVw6T1jVUgyP3h6NsWPplsx_frFRG5TMKVsRKgFvWOULmor-99CnwKOiEl3JtTBW9_whuDBHJ9DER-VRY6kvU0xbMoNBE8CEJ-klT6BDfaLCQDY-PrMsTBdqJn7KMfBiHjzUjvhfA0OUBSSzwX_uLPKoZKp9HAyXr368xUXTmBJIQNFroV8AVw_Aa2HJZsh7TayxnLtPr-PA97Z4RG0ZLHJMzYoNTcUclNGo3Eq_Sr6moU3oeqUQ3fzbokZDI8aiq4FcK8mxvTA2mMyXTCfZpszUOw3bhTeLeTkbP0.Z1vayhJffy63Y4Mi3hdjHbEIvu9uVWcgEhtnwtLUxZY\u0026dib_tag=se\u0026keywords=picture%2Bframe%2Bhanging%2Bkit%2Bleveler%2Btool\u0026qid=1770505093\u0026sr=8-1","price":"$32.92","error":""},{"name":"LEVOIT Air Purifiers for Bedroom Home Dorm, 3-in-1 Filter Cleaner With Aroma Pad (Core Mini-P)","image_path":"https://m.media-amazon.com/images/I/71wGv7Fh2AL._AC_UY218_.jpg","provider_url":"https://www.amazon.com/LEVOIT-Purifiers-Freshener-Core-Mini/dp/B09GTRVJQM/ref=sr_1_1?dib=eyJ2IjoiMSJ9.66Bvqtkv21NZp4opfcsoRO9JcSm6y_Y_Gj7LHndUQcwUYJm4o25hjKNdofuBZBzlWKaFqNN9-zv8JtMPrklY7ZEIG-eWwzS5cng1a7MXMsQTVipbepNan5mGWQS0Jpms20jqNF6RDKf3mZfNttxx0UhIydF28XjXyiRGuff-rM_TV9wgSVIyTQtht_sUDCwSbijEkJIAxf1C8_R0CP0E16GJkYJxv2WiTLngmnP1lIc.x9_RqL71_9laVRdtbkO--YWQzEXS-IL-FwXlFlXmCzg\u0026dib_tag=se\u0026keywords=air%2Bpurifier%2Bfor%2Bbedroom%2BHEPA%2Bcompact\u0026qid=1770505092\u0026sr=8-1","price":"$41.99","error":""}]}`
		writeResponse(w, http.StatusOK, mockResult)
		return
	// }


	orders, err := database.GetAllOrder(r.Context())
	if err != nil {
		panic(":(")
	}

	buf, err := json.Marshal(&orders)
	ordersStringified := string(buf)
	var startPrompt = fmt.Sprintf(`You are Precog, a proactive shopping assistant. Your role is to help the user research, compare, and choose products, find the best deals, suggest alternatives, and warn about stock, prices, compatibility, or scams. Use all available tools and context to anticipate needs, provide reminders, and give concise, actionable advice tailored to the user’s preferences, budget, and trends. All the data must come from one of defined providers, so if you have an idea of what you want to recommend, YOU MUST RETURN REAL RESULTS FROM A PROVIDER.

	Here is some user order history from the user. You should use this data to creatively give recommendations.
	Order History: %s

	Here is an example of how you can use order history to give useful recommendations:
		* If the order history shows many orders with household items, perform an amazon search query for a household item they haven't yet bought (ie. A night light), and provide the Amazon search results as structured JSON data from the search_amazon tool call.

	Use the tools named search_amazon to search for Amazon products to give Amazon recommendations. YOUR RESPONSE SHOULD BE A LIST WHERE EACH ITEM FOLLOWS THIS SCHEMA:
		{
			name: string;
			image_path: string;
			provider_url: string;
			price: string;
		} or {
			error: string
		}

DO NOT RESPOND WITH ANY OTHER TEXT THAN THE JSON SCHEMA. Output only a JSON array.`, ordersStringified)

	var userPrompt = `Based on all the data you have about me, provide interesting recommendations I might benefit from. Be ambigious and creative. Do not tell me you cannot answer the question, you have the context and capabilities. These can include suggestions, warnings, reminders, or optimizations across any aspect of my shopping habits, preferences, budget, or trends. Highlight what actions or changes could improve my choices or save me time or money.`

	response, err := llm.Call([]llm.Message{
		llm.NewMessage(startPrompt, llm.SystemMessage),
		llm.NewMessage(userPrompt, llm.UserMessage),
	}, build_llm_tools(r.Context()))

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type Recommendation struct {
		Name        string `json:"name"`
		ImagePath   string `json:"image_path"`
		ProviderURL string `json:"provider_url"`
		Price       string `json:"price"`
		Error       string `json:"error"`
	}

	var recommendations []Recommendation
	if err := json.Unmarshal([]byte(response), &recommendations); err != nil {
		log.Printf("Recommendations LLM returned invalid JSON: %v", err)
		writeError(w, http.StatusInternalServerError, "LLM returned invalid JSON")
		return
	}

	writeResponse(w, http.StatusOK, recommendations)
}
