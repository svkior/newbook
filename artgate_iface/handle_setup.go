package main
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"fmt"
	"strconv"
)


func HandleSetupEthEdit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "setup/editip", map[string]interface{}{
		"Setup": globalSetup,
	})
}

func HandleSetupEthUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ipAddr := r.FormValue("ipaddr")
	ipMask := r.FormValue("ipmask")
	ipGw   := r.FormValue("ipgw")
	ethMac := r.FormValue("macs")


	log.Printf("MACS: %s", ethMac)

	err := globalSetup.UpdateIpAddr(ipAddr)
	if err != nil {
		if IsValidationError(err) {
			RenderTemplate(w, r, "setup/editip", map[string]interface{}{
				"Error": err.Error(),
				"Setup":  globalSetup,
			})
			return
		}
		panic(err)
	}

	err = globalSetup.UpdateIpMask(ipMask)
	if err != nil {
		if IsValidationError(err) {
			RenderTemplate(w, r, "setup/editip", map[string]interface{}{
				"Error": err.Error(),
				"Setup":  globalSetup,
			})
			return
		}
		panic(err)
	}

	err = globalSetup.UpdateIpGateway(ipGw)
	if err != nil {
		if IsValidationError(err){
			RenderTemplate(w, r, "setup/editip", map[string]interface{}{
				"Error": err.Error(),
				"Setup": globalSetup,
			})
			return
		}
		panic(err)
	}

	err = globalSetup.UpdateMac(ethMac)
	if err != nil {
		if IsValidationError(err){
			RenderTemplate(w, r, "setup/editip", map[string]interface{}{
				"Error": err.Error(),
				"Setup": globalSetup,
			})
			return
		}
		panic(err)
	}

	http.Redirect(w, r, "/?flash=Ip+Addr+updated", http.StatusFound)
}

func HandleSetupArtnetEdit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "setup/editartnet", map[string]interface{}{
		"Setup": globalSetup,
	})
}

func HandleSetupArtnetUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	numInputs := r.FormValue("artinputs")

	log.Println(numInputs)

	err := globalSetup.UpdateArtNetInputs(numInputs)
	if err != nil {
		if IsValidationError(err){
			RenderTemplate(w, r, "setup/editartnet", map[string]interface{}{
				"Error": err.Error(),
				"Setup": globalSetup,
			})
			return
		}
		panic(err)
	}

	for idx:=0; idx < globalSetup.ArtnetInputs;idx++{

		_, err := strconv.Atoi(r.FormValue(fmt.Sprintf("tag%dOption",idx)))
		if err != nil {
			globalSetup.DisableArtnetIn(idx)
		} else {
			log.Printf("Idx: %d, Sw: %s\n",idx, r.FormValue(fmt.Sprintf("tag%d",idx)))
			globalSetup.EnableArtnetIn(idx)
			err = globalSetup.UpdateArtNetInUniverse(idx, r.FormValue(fmt.Sprintf("tag%d",idx)))
			if err != nil {
				if IsValidationError(err){
					RenderTemplate(w, r, "setup/editartnet", map[string]interface{}{
						"Error": err.Error(),
						"Setup": globalSetup,
					})
					return
				}
				panic(err)
			}
		}
	}

	http.Redirect(w, r, "/?flash=Artnet+updated", http.StatusFound)
}

func HandleSetupArtnetOutEdit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "setup/editartnetout", map[string]interface{}{
		"Setup": globalSetup,
	})
}

func HandleSetupArtnetOutUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	numOutputs := r.FormValue("artinputs")

	//log.Println(numOutputs)
	err := globalSetup.UpdateArtNetOutputs(numOutputs)
	if err != nil {
		if IsValidationError(err){
			RenderTemplate(w, r, "setup/editartnetout", map[string]interface{}{
				"Error": err.Error(),
				"Setup": globalSetup,
			})
			return
		}
		panic(err)
	}
	for idx:=0; idx < globalSetup.ArtnetOutputs;idx++{

		_, err := strconv.Atoi(r.FormValue(fmt.Sprintf("tag%dOption",idx)))
		if err != nil {
			globalSetup.DisableArtnetOut(idx)
		} else {
			log.Printf("Idx: %d, Sw: %s\n",idx, r.FormValue(fmt.Sprintf("tag%d",idx)))
			globalSetup.EnableArtnetOut(idx)
			err = globalSetup.UpdateArtNetOutUniverse(idx, r.FormValue(fmt.Sprintf("tag%d",idx)))
			if err != nil {
				if IsValidationError(err){
					RenderTemplate(w, r, "setup/editartnetout", map[string]interface{}{
						"Error": err.Error(),
						"Setup": globalSetup,
					})
					return
				}
				panic(err)
			}
		}
	}

	http.Redirect(w, r, "/?flash=Artnet+updated", http.StatusFound)
}
