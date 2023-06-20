package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql" //bliblioothèque des drivers permettant la liaison et l'interaction vers le serveur mysql
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	id_utilisateurs int
	id_messages     int
	id_compte       int    // id de compte de l'utilisateur connecté
	nom_session     string // nom  de l'utilisateur connecté
	prenom_session  string // prénom  de l'utilisateur connecté
	id              int
	dateheure       string
	contenu         string
	visibilite      int
	image           string
	update_at       string
	date_creation   sql.NullString
	id_topic        int
	mail            string // adresse mail
	motdepasse      string
	Sujet           string
	messages_length int
)

type value struct {
	Result          string // variable qui sera affichée si il n'existe pas de page  topic avec l'id donnée
	Dateheure_topic string // heure de création du topic
	Sujet           string // Titre du sujet
	Connecté        string // variable  des données de connection (nom et prénom)
	Data            string //
	Connexionresult string
}

var login = template.Must(template.ParseFiles("HTML/Login.html"))
var register = template.Must(template.ParseFiles("HTML/Register.html"))
var index = template.Must(template.ParseFiles("HTML/index.html"))
var topic = template.Must(template.ParseFiles("HTML/Topics.html"))
var reply = template.Must(template.ParseFiles("HTML/ReponseTopic.html")) // chacune des variables utilise les différent fichiers présents dans le dossier HTML
var cgu = template.Must(template.ParseFiles("HTML/CGU.html"))
var policy = template.Must(template.ParseFiles("HTML/Politique.html"))
var about = template.Must(template.ParseFiles("HTML/Equipe.html"))
var contactus = template.Must(template.ParseFiles("HTML/Contact.html"))
var like = template.Must(template.ParseFiles("HTML/Topics.html"))
var dislike = template.Must(template.ParseFiles("HTML/Topics.html"))

func main() {
	fmt.Println(id_compte)
	styleServer := http.FileServer(http.Dir("CSS"))              // indique au serveur ou se trouve le fichier CSS
	http.Handle("/CSS/", http.StripPrefix("/CSS/", styleServer)) // indication du fichier contenant le dossier

	http.HandleFunc("/login", HttpHandlerLogin)         // route menant vers la page  de connexion
	http.HandleFunc("/register", HttpHandlerRegister)   // route menant vers la page d'inscription
	http.HandleFunc("/", httpHandlerIndex)              // route menant vers la page d'acceuil
	http.HandleFunc("/topic/", httpHandlerTopic)        //route menant vers une page de topic avec une id donnée
	http.HandleFunc("/logout", HttpLogoutHandler)       //route menant vers la page de déconnection
	http.HandleFunc("/answer/", HttpAnswerHandler)      //route menant vers la page permettant dev répondre à un message
	http.HandleFunc("/policy", HttpPolicyHandler)       //route  menant vers la page de politique  du forum
	http.HandleFunc("/cgu", HttpCGUHandler)             //route menant vers la page de cgu
	http.HandleFunc("/about", HttpAboutHandler)         // route menant vers la page à propos
	http.HandleFunc("/contactus", HttpContactUsHandler) // route menant vers  la page de contact
	http.HandleFunc("/like/", HttpLikeHandler)          //route menant vers la page permettant de liker
	http.HandleFunc("/dislike/", HttpDislikeHandler)    // route menant  la page permettant de disliker

	http.ListenAndServe(":8080", nil) //lancement du serveur sur le port 80

	dbHost := "localhost"
	dbPort := "3306"  //adress url du serveur
	dbUser := "root"  // nom utilisateur
	dbPass := ""      //mot de passe
	dbName := "forum" //nom_bdd

	// Chaîne de connexion
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName) // association des données de connection

	// Connexion à la base de données
	db, err := sql.Open("mysql", dbURI) // accès à la base  de données
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database is reachable", db, dbURI) // connection réussie

	select1, err := db.Query("SELECT * FROM message")

	if err != nil {
		panic(err)
	}

	for select1.Next() {

		err1 := select1.Scan(&id,
			&dateheure,
			&contenu,
			&visibilite,
			&image,
			&update_at,
			&date_creation,
			&id_topic,
			&id_utilisateurs) //&id_utilisateurs, &id_messages

		if err1 != nil {

			panic(err1.Error())
		}

		fmt.Println(id_utilisateurs, id_utilisateurs, contenu)
		fmt.Println(select1)
	}

	defer select1.Close()

	select2, err := db.Query("SELECT * FROM message") //aimer

	if err != nil {
		panic(err)
	}

	for select2.Next() {

		err1 := select1.Scan(&id,
			&dateheure,
			&contenu,
			&visibilite,
			&image,
			&update_at,
			&date_creation,
			&id_topic,
			&id_utilisateurs) //&id_utilisateurs, &id_messages

		if err1 != nil {

			panic(err1.Error())
		}

		fmt.Println(id_utilisateurs, id_utilisateurs, contenu)
		fmt.Println(select1)
	}

	defer select2.Close()

}

func HttpHandlerLogin(w http.ResponseWriter, r *http.Request) { // fonction de  la page connection

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root" // idem que précèdement
	dbPass := ""
	dbName := "test" //nom_bdd

	// Chaîne de connexion
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Connexion à la base de données
	db1, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	var checkusername = r.FormValue("mail")     // valeur du  champs ayant le nom mail
	var checkpassword = r.FormValue("password") // valeur du  champs ayant le nom password

	search, err0 := db1.Query("SELECT mail FROM utilisateurs WHERE mail = ?  ", checkusername) //  on vérifie si il existe  bien une email comme celle rentrée par l'utilisateur

	if err0 != nil {
		panic(err0.Error())
	}

	for search.Next() { // on prépare à scanner les information renvoyées par la base de données
		err1 := search.Scan(&mail) // on scanne la valeur renvoyée
		if err1 != nil {           //gestion d'erreur
			fmt.Println("Error")
			panic(err1.Error())
		}
		fmt.Println(mail)
	}
	fmt.Println(motdepasse, mail)
	mailtocompare := mail
	search.Close() // fermeture de la colonne de recherche
	fmt.Println(checkpassword, "test0")
	hash1 := sha256.Sum256([]byte(checkpassword)) // Hachage de la valeur entrée par l'utilisateur pour comparaison
	hashString := hex.EncodeToString(hash1[:])    //on convertit la valeur obtenue  en string

	checkpassword = hashString

	search1, err2 := db1.Query("SELECT motdepasse,id,prenom,nom FROM utilisateurs WHERE motdepasse = ? AND mail= ? ", checkpassword, checkusername) // recherche d'un utilisateur et de ses données (motdepasse,id,prenom,nom) dans la base de donnée  ayant le nom et l'adresse mail renseignés

	if err2 != nil {
		panic(err2.Error())
	}
	var id_util int
	var nom_util string
	var prenom_util string
	for search1.Next() {
		err3 := search1.Scan(&motdepasse, &id_util, &prenom_util, &nom_util)
		if err3 != nil {
			fmt.Println("Error")
			panic(err3.Error())
		}

	}
	id_compte = id_util
	nom_session = prenom_util
	prenom_session = nom_util
	passwordtocompare := motdepasse
	search1.Close()
	var Result string // message qui sera renvoyé  avec {{.Result}} et Execute() et sprintf
	var invalidcredits bool

	if checkpassword != passwordtocompare { // si les mots de passe ne correspondent pas
		log.Println("erreur le mot de passe entré est incorrect")
		invalidcredits = true
	}

	if checkusername != mailtocompare { // si les adresses ne correspondent pas
		log.Println("erreur l'adresse mail entrée est incorrecte")
		invalidcredits = true

	}

	if invalidcredits == true { //si l'une des ou les deux valeurs renseignées sont incorrectes
		log.Println("Echec de la connexion")
		Result = "l'adresse mail ou le mot de passe est incorrect" // ""
	} else {
		log.Println("Réussite de la connection")
		Result = "Connexion réussie"

		http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)

	}
	data := value{
		Result: fmt.Sprintf(Result)}
	login.Execute(w, data)

}

func HttpHandlerRegister(w http.ResponseWriter, r *http.Request) { // fonction qui gère la page d'inscription

	var nametoadd string     // nom à ajouter
	var surnametoadd string  // prénom à ajouter
	var emailtoadd string    // email à ajouter
	var passwordtoadd string // mot de passe à ajouter

	nametoadd = r.FormValue("username")     // valeur du champ nom renseigné par l'utilisateur
	surnametoadd = r.FormValue("surname")   // valeur du champ prénom renseigné par l'utilisateur
	emailtoadd = r.FormValue("email")       // valeur du champ email renseigné par l'utilisateur
	passwordtoadd = r.FormValue("password") // valeur du mot de passe nom renseigné par l'utilisateur
	fmt.Println(nametoadd, surnametoadd, emailtoadd, passwordtoadd)

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root" // accès à la base de donnée
	dbPass := ""
	dbName := "test"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Connexion à la base de données
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database is reachable", db, dbURI)

	select1, err := db.Prepare("INSERT INTO utilisateurs  (id_roles,nom,prenom,mail,motdepasse) VALUES(?,?,?,?,?);") // on formate notre requete sql ou chaque ? sera remplacé par une variable dans le meme ordre

	if err != nil {
		panic(err)
	}

	defer select1.Close()
	var result string
	if nametoadd != "" && surnametoadd != "" && emailtoadd != "" && passwordtoadd != "" { //avant d'enregister ces valeurs on s'assure qu'elles sont pas "vides"

		hash := sha256.Sum256([]byte(passwordtoadd)) // on Hache le mot de passe avec le hachage sa256
		hashString := hex.EncodeToString(hash[:])    // on convertit en string
		passwordtoadd = hashString

		_, registernewuser := select1.Exec("1", ""+nametoadd+"", ""+surnametoadd+"", ""+emailtoadd+"", ""+passwordtoadd+"") // enregitrement des données  dans Base de Données
		if registernewuser != nil {                                                                                         // gestion d'erreur  qui va nous permettre d'empecher au serveur de crash si l'adresse mail est déja prise (unique)

			fmt.Println(registernewuser.Error())
			test := string(registernewuser.Error()) //conversion de l'erreur en  string
			errorcode := test[5:10]                 //récuperation de la slice qui contient le nom de l'erreur
			fmt.Println(errorcode)

			if errorcode == " 1062" { // si le code équivaut au code  d'erreur s'affichant lorsqu'une adresse mail est déja prise

				fmt.Println("erreur l'adresse mail est déja prise ")
				result = "erreur l'adresse mail est déja prise " // message renvoyé à l'utilisateur

			} else { // autre erreur que l'on attrape pas
				log.Panic("erreur  fatale : ", registernewuser)
			}

		}

		select2, err := db.Query("SELECT id FROM utilisateurs WHERE nom = ?", nametoadd) //récupération de l'id de l'utilisateur nouvellement créé

		if err != nil {
			panic(err)
		}

		var id_user int
		select2.Next()

		err1 := select2.Scan(&id_user) //&id_utilisateurs, &id_messages

		if err1 != nil {

			panic(err1.Error())
		}

		defer select2.Close()
		id_compte = id_user //  l'id du compte connecté l'id récupérée
		fmt.Println(id_user, "USER_ID", id_compte)
		nom_session = nametoadd       // idem pour le nom du compte
		prenom_session = surnametoadd //  idem pour le prénom du compte

		http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther) // rédirige vers la page d'acceuil l'utilisateur connecté

		//log.Panic(registernewuser)

	} else { // si il manques des donnée
		fmt.Println("des données sont manquantes (erreur416554665)")
		result = "Veillez compléter tous les champs" //message renvoyé à l'utilisateur
	}
	data := value{
		Result: fmt.Sprintf(result),
	}
	register.Execute(w, data)
}

func httpHandlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EST VIDE ?", id_compte)

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root" //nouvel accès à la base de données
	dbPass := ""
	dbName := "test"
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database is reachable", db, dbURI)

	select0, err := db.Query("SELECT COUNT(*) FROM topic") // récuperation du nombre de topic

	if err != nil {
		panic(err)
	}
	var nb int
	for select0.Next() {

		err1 := select0.Scan(&nb)

		if err1 != nil {

			panic(err1.Error())
		}

		fmt.Println(nb)

	}

	defer select0.Close()

	select1, err := db.Query("SELECT * FROM topic") // récupération de toute les information de tout les topics

	if err != nil {
		panic(err)
	}
	var lastdate string
	var lastesujet string

	var data string = "<div class='ParentTopics'>" // ajout  du début de la div  Parent Topic toutes les div contenant les topic seront contenues dans cette div
	for i := 0; i < nb; i++ {                      // pour le nombre  de topic (peut remplacer for next() )
		select1.Next()
		err1 := select1.Scan(&id, //on récupère l'heure de création,l'id et le sujet
			&dateheure,
			&Sujet)

		if err1 != nil {

			panic(err1.Error())
		}

		fmt.Println(id, dateheure, Sujet, i)
		var idstring = strconv.Itoa(id) // convertion en string
		lastdate = dateheure            // heure de creation du topic
		lastesujet = Sujet              // nom du sujet
		// génération du html avec toute les information des topic (concaténation)
		data = data + "<a  class='LiensTopics' href='http://localhost:8080/topic/" + idstring + "'><div class='Topics'><div class='DatesHeuresTopics'><p class='Auteur'></p><p class='Heures'>Posté le " + dateheure + "</p></div><div class='ParentTitreTopics'><h1 class='TitreTopics'>" + Sujet + "</h1></div></div></a>"

	}
	defer select1.Close()
	fmt.Println(id, dateheure, Sujet, lastdate, lastesujet, "last")
	data = data + "</div>" //fermeture de  la div Parent Topic

	select2, err := db.Query("SELECT COUNT(*) FROM message") // calcul du nombre de message

	if err != nil {
		panic(err)
	}
	var len int //nombre des messages

	select2.Next()
	err1 := select2.Scan(&messages_length)

	if err1 != nil {

		panic(err1.Error())
	}

	fmt.Println(messages_length)

	defer select2.Close()
	len = messages_length - 1
	fmt.Println(len, messages_length)

	select3, err := db.Query("SELECT contenu,id_utilisateurs,id_topic,dateheure FROM `message`  LIMIT 1 OFFSET ?", len) // on récupère le dernier message (le plus récent)

	if err != nil {
		panic(err)
	}
	var contenu_msg string  //contenu du message (dernier message)
	var id_utilisateurs int // id utilisateur
	var id_topic int        // id du topic

	select3.Next()
	err2 := select3.Scan(&contenu_msg, &id_utilisateurs, &id_topic, &dateheure) //scan du résultat

	if err2 != nil {

		panic(err2.Error())
	}

	fmt.Println(contenu_msg, id_utilisateurs, id_topic, dateheure)

	defer select3.Close()

	select4, err := db.Query("SELECT nom,prenom FROM `utilisateurs`  WHERE id = ?", id_utilisateurs) //on récpère le nom et prénom à partir de l'id utilisateur

	if err != nil {
		panic(err)
	}
	var nom string
	var prenom string

	select4.Next()

	err3 := select4.Scan(&nom, &prenom)

	if err3 != nil {

		panic(err3.Error())
	}

	fmt.Println(nom, prenom)

	if id_compte == 0 { // si l'id est celle par défaut
		nom_session = "non"
		prenom_session = "connecté"
	}

	defer select4.Close()

	data = data + "<div class='ParentDernierMessages'>" // ajout de  la div ParentDernierMessages qui va stocker le dernier message
	//data = ""
	if strings.Contains(contenu_msg, "Ø") { // si le message est une réponse
		var temp string                    //contient une valeur de la string avant son ajout dans le tableau
		var tab []string                   //tableau qui va contenir tout les info
		for _, part := range contenu_msg { // on parcours toute la string
			if string(part) != "Ø" { // si la partie de string est différente du Ø séparateur
				temp = temp + string(part) // on concatène  a temp
				fmt.Println("PIECE", string(part))
			}
			if string(part) == "Ø" { // si la valeur de la string  est un séparateur

				fmt.Println(temp)
				tab = append(tab, string(temp)) // on ajoute la valeur au tableau
				temp = ""
			}

		}
		//nos valeurs sont maintenant séparées
		contenu_msg = tab[1] // le contenu de la réponse se trouve à l'index 1
	}

	// génération du html avec toute les information du dernier message  (concaténation) ainsi qu'une div pour rédiger un nnouveau topic
	data = data + "<div class='Topics' id='latestmessage'> <h1 id='lastmsg'>  Dernier message:</h1><div class='DatesHeuresTopics'><p class='Auteur'>Publié par : " + prenom + " " + nom + "</p><p class='Heures'> posté à :" + dateheure + "</p></div><a class='LiensTopic' href='http://localhost:8080/topic/" + strconv.Itoa(id_topic) + "' style='text-decoration:none;color:white;'><div class='ParentTitreTopics'><h1 class='TitreTopics'>" + prenom + " " + nom + "</h1><p class='MessagesTopics'>" + contenu_msg + "</p></div></a><div class='Réaction'><div class='Répondre'><a class='Réponse' href='../page/ReponseTopic.html'><form method='POST' action='http://localhost:8080/answer/" + strconv.Itoa(id_messages) + "-" + nom_session + "-" + prenom_session + "-" + prenom + "-" + nom + "-" + strconv.Itoa(id) + "-" + "''><button  type='submit'class='FondBlancRéaction'>Répondre</button></form></a></div><img class='Like' src='https://raw.githubusercontent.com/deverror6068/images-forum/main/Like.png' alt='Like'><img class='Dislike' src='https://raw.githubusercontent.com/deverror6068/images-forum/main/Dislike.png' alt='Dislike'></div></div></div>"

	data = data + "<div class='inputtext' id='newtopic'><form><h1>Nouveau  Topic</h1><input name='newmsg' type='text' placeholder='entrez le nom du sujet' required><button class='EnvoieNouveauxMessages' type='submit'>Envoyer message</submit></form></div>"

	data = data + "</div></div>" // fermeture des viv

	var msgcontent = r.FormValue("newmsg") // contenu du message entré par l'utilisateur

	if msgcontent == "" { // si le message est vide
		log.Println("message vide ")
	} else {

		if id_compte == 0 { // si l'utilisateur n'est pas connecté

			fmt.Println("Vous devez etre connecté !")

		} else {

			var time = time.Now().Format("2006-01-02 15:04") // on récupère le temps actuel et  on formate le temps au format aaaa:mm:jj:hh:mm
			select1, err := db.Prepare("INSERT INTO topic  (dateheure,sujet) VALUES(?,?);")

			if err != nil {
				panic(err)
			}

			_, registernewuser := select1.Exec(""+time+"", ""+msgcontent+"") //ajout du nom et de la date  à la BD
			if registernewuser != nil {
				panic(registernewuser)
			}

			select5, err := db.Query("SELECT id FROM `topic`  WHERE sujet = ?", msgcontent) // récupère l'id pour la redirection

			if err != nil {
				panic(err)
			}
			var id_msg string

			select5.Next()

			err4 := select5.Scan(&id_msg)

			if err4 != nil {

				panic(err3.Error())
			}

			defer select4.Close()

			http.Redirect(w, r, "http://localhost:8080/topic/"+id_msg, http.StatusSeeOther)
			msgcontent = ""
		}
	}
	token := value{Connecté: fmt.Sprintf(nom_session + " " + prenom_session),
		Data: data}

	index.Execute(w, token)
	w.Write([]byte(data)) // génération du html

}
func httpHandlerTopic(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/topic/"):] // récupère l'adresse url  à partir du dernier / de topic
	id, err := strconv.Atoi(idStr)       // on transforme les chiffre qui sont en string en entier
	if err != nil {                      // si la conversion n'est pas possible
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	log.Println(id, "Id du Topic")

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "test"
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("DBNAME", dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database is reachable", db, dbURI)

	select4, err := db.Query("SELECT sujet FROM `topic`  WHERE id = ?", id) // verifie si le topic a un sujet (si il existe)

	if err != nil {
		panic(err)
	}
	var id_topic string
	var Result string

	for select4.Next() {

		err3 := select4.Scan(&id_topic)

		if err3 != nil {

			panic(err3.Error())
		}

		fmt.Println(id_topic)
	}
	defer select4.Close()

	if id_topic == "" { // si le topic n'a pas de sujet il n'existe pas

		fmt.Println("ECHEC !!!!!")
		Result = "Aucun resultat"
		data := value{
			Result: fmt.Sprintf(Result),
		}
		topic.Execute(w, data)

	} else {

		log.Println("correspondance trouvée ")
		//on récupère toute les données qui utiles à afficher
		select1, err := db.Query("SELECT contenu,message.id,dateheure,nom,prenom,contenu FROM `utilisateurs` JOIN message ON( id_utilisateurs = utilisateurs.id AND id_topic = ?)", id) //jointure premettant de récupérer toute ces données

		if err != nil {
			panic(err)
		}
		var contenu string
		var dateheure string
		var nom string
		var prenom string
		var id_messages int
		var data_msg string = ""
		//var result string
		for select1.Next() { // pour le nombre de messages

			err1 := select1.Scan(&contenu, &id_messages,
				&dateheure,
				&prenom, &nom, &contenu)

			if err1 != nil {

				panic(err1.Error())
			}

			fmt.Println(id, dateheure, Sujet)

			fmt.Println(id, dateheure, contenu, nom, prenom, id_messages, "least")

			select5, err := db.Query("SELECT COUNT(id_utilisateurs) FROM `appreciation_messages` WHERE id_messages = ? AND etat_like = 1;", id_messages) // on récupère le nombre de like du message

			if err != nil {
				panic(err)
			}
			var likecount string // nombre de like

			select5.Next()

			err4 := select5.Scan(&likecount)

			if err4 != nil {

				panic(err4.Error())
			}

			fmt.Println(likecount)

			defer select5.Close()

			select6, err := db.Query("SELECT COUNT(id_utilisateurs) FROM `appreciation_messages` WHERE id_messages = ? AND etat_like = 0;", id_messages) // on récupère le nombre de dislike

			if err != nil {
				panic(err)
			}
			var dislikecount string // nombre de dislike

			select6.Next()

			err5 := select6.Scan(&dislikecount)

			if err5 != nil {

				panic(err5.Error())
			}

			fmt.Println(dislikecount)

			defer select6.Close()

			data_msg = data_msg + " <div class='ParentTopics'>" //div qui   stocke chaque message
			if strings.Contains(contenu, "Ø") {                 // si le messages est une reponse
				var temp string
				var tab []string
				for _, part := range contenu { // on sépare les données contenues dans l'url

					if string(part) != "Ø" {
						temp = temp + string(part)
					}
					if string(part) == "Ø" {

						fmt.Println(temp)
						tab = append(tab, string(temp))
						temp = ""
					}

				}
				select2, err := db.Query("SELECT id FROM `utilisateurs`  WHERE nom = ?", tab[3]) //on récupère l'id de de l'utilisateur auquel le l'utilisateur connecté réponds

				if err != nil {
					panic(err)
				}
				var idoriginal int

				select2.Next()

				err3 := select2.Scan(&idoriginal)

				if err3 != nil {

					panic(err3.Error())
				}
				defer select2.Close()

				select3, err := db.Query("SELECT dateheure FROM `message`  WHERE id_utilisateurs = ?", idoriginal) // on récupère l'heure de création du message

				if err != nil {
					panic(err)
				}
				var dateoriginal string

				select3.Next()

				err4 := select3.Scan(&dateoriginal)

				if err4 != nil {

					panic(err3.Error())
				}
				defer select4.Close()

				var contenuorigine string = tab[0]
				var reponse string = tab[1]

				// génération du html avec toute les information du  message  (réponse)
				data_msg = data_msg + "<div class='Topics' name = 'feedtopic' id='" + strconv.Itoa(id_messages) + "' style='height: 55% ;'><div class='DatesHeuresTopics'><h1 class='TitreTopics'>Réponse à " + tab[2] + " " + tab[3] + "</h1><p class='Auteur'> " + tab[2] + " " + tab[3] + " A dit :</p></div><div class='ParentTitreTopics'><p class='MessagesTopics'>" + contenuorigine + " </p><p class='Heures'>Posté le : " + dateoriginal + " </p></div><div class='DatesHeuresTopics'></div><div class='ParentTitreTopics'><h1 class='TitreTopics'>" + prenom + " " + nom + "</h1><p class='MessagesTopics'>" + reponse + "</p>  <p class='Heures'>Posté le : " + dateheure + "</p></div><div class='Réaction'><div class='Répondre'><form method='POST' action='http://localhost:8080/answer/" + strconv.Itoa(id_messages) + "-" + nom_session + "-" + prenom_session + "-" + prenom + "-" + nom + "-" + strconv.Itoa(id) + "-" + "''><button  type='submit'class='FondBlancRéaction'>Répondre</button></form><p>" + likecount + "</p><a href='http://localhost:8080/like/" + strconv.Itoa(id_messages) + "-" + nom_session + "-" + prenom_session + "-" + strconv.Itoa(id) + "-" + "'></div><img class='Like' src='https://raw.githubusercontent.com/deverror6068/images-forum/main/Like.png' alt='Like'></a><p>" + dislikecount + "</p><a href='http://localhost:8080/dislike/" + strconv.Itoa(id_messages) + "-" + nom_session + "-" + prenom_session + "-" + strconv.Itoa(id) + "-" + "'><img class='Dislike' src='https://raw.githubusercontent.com/deverror6068/images-forum/main/Dislike.png' alt='Dislike'></a></div></div>"
			} else { // génération du html avec toute les information du  message ("message normal")
				data_msg = data_msg + "<div class='Topics'  name = 'feedtopic' id='" + strconv.Itoa(id_messages) + "'> <div class='DatesHeuresTopics'><p class='Auteur'>Publié par : " + prenom + " " + nom + "</p><p class='Heures'> posté à :" + dateheure + "</p></div><div class='ParentTitreTopics'><h1 class='TitreTopics'>" + prenom + " " + nom + "</h1><p class='MessagesTopics'>" + contenu + "</p></div><div class='Réaction'><div class='Répondre'><form method='POST' action='http://localhost:8080/answer/" + strconv.Itoa(id_messages) + "-" + nom_session + "-" + prenom_session + "-" + prenom + "-" + nom + "-" + strconv.Itoa(id) + "-" + "''><button  type='submit'class='FondBlancRéaction'>Répondre</button></form></div><p>" + likecount + "</p><a href='http://localhost:8080/like/" + strconv.Itoa(id_messages) + "-" + nom_session + "-" + prenom_session + "-" + strconv.Itoa(id) + "-" + "'><img class='Like' src='https://raw.githubusercontent.com/deverror6068/images-forum/main/Like.png' alt='Like'></a><p>" + dislikecount + "</p><a href='http://localhost:8080/dislike/" + strconv.Itoa(id_messages) + "-" + nom_session + "-" + prenom_session + "-" + strconv.Itoa(id) + "-" + "'><img class='Dislike' src='https://raw.githubusercontent.com/deverror6068/images-forum/main/Dislike.png' alt='Dislike'></a></div></div></div>"
			}
		}
		defer select1.Close()

		var msgcontent = r.FormValue("newmsg") //géneration d'un pour écrire un nouveau message
		data_msg = data_msg + "<div class='inputtext' id = 'newfeed'><form><h1>Nouveau Messages</h1><input name='newmsg' type='text' required><button class='EnvoieNouveauxMessages' type='submit'>Envoyer message</submit></form></div>"
		fmt.Println(msgcontent, "Contenu du message ")

		if msgcontent == "" {
			fmt.Println("Le message est vide")
			topic.Execute(w, data_msg)
			w.Write([]byte(data_msg)) //génération du html
		} else {
			fmt.Println("Id utilisateur connecté", id_compte)
			if id_compte == 0 { // si l'utilisateur ne s'est pas connecté
				fmt.Println("ERREUR VOUS DEVEZ ETRE CONNECTE POUR POUVOIR ECRIRE UN MESSAGE")
				data_msg = data_msg + "<p>erreur vous devez etre connecte pour pouvoir ecrire un message</p>"

			} else {
				var time = time.Now().Format("2006-01-02 15:04") //formatage du temps
				select1, err := db.Prepare("INSERT INTO message  (contenu,dateheure,id_topic,id_utilisateurs) VALUES(?,?,?,?);")

				if err != nil {
					panic(err)
				}

				_, registernewuser := select1.Exec(""+msgcontent+"", ""+time+"", ""+strconv.Itoa(id)+"", id_compte) // on insère à la BD toute les info du messages
				if registernewuser != nil {
					panic(registernewuser)
				}

				http.Redirect(w, r, "http://localhost:8080/topic/"+strconv.Itoa(id)+"#"+strconv.Itoa(id_messages), http.StatusSeeOther) // on redirige vers le message nouvellement crée
				msgcontent = ""
			}

			topic.Execute(w, data_msg)
			w.Write([]byte(data_msg))
		}

	}
}

func HttpLogoutHandler(w http.ResponseWriter, r *http.Request) { // fonction de la page de déconnexion

	id_compte = 0 // l'id de connexion devient celui par défaut (non connecté)
	fmt.Println("id du compte connecté", id_compte)
	data := value{}
	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther) // redirection vers la page d'acceuil
	index.Execute(w, data)

}

func HttpAnswerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("URLPATH", r.URL.Path[:], len(r.URL.Path[:]))
	var cheminentier = r.URL.Path[:] // on récupère le chemin url entier

	fmt.Println("NOM SESSION", nom_session, prenom_session)

	fmt.Println(id, "ANSWER")
	fmt.Println(len(cheminentier), "FULL PATH")

	var tab []string
	var temp string

	for _, part := range cheminentier { // on sépare chaque donnée presente dans l'url
		if string(part) != "-" { // si la partie de string n'est pas séparateur
			temp = temp + string(part) // on l'ajoute a temp
			fmt.Println("PIECE", string(part))
		}
		if temp == "/answer/" { // si temp  vaut la première partie de l'url on l'ignore
			temp = ""
		}
		if string(part) == "-" { // si  la partie de string est un séparateur

			fmt.Println(temp)
			tab = append(tab, string(temp)) // on ajoute la valeur au tableau
			temp = ""
		}

	}
	if id_compte == 0 { // si l'utilisateur n'est pas connecté
		nom_session = "non"
		prenom_session = "connecté"
		http.Redirect(w, r, "http://localhost:8080/topic/"+tab[5], http.StatusSeeOther) // on le redirige vers le message
	}

	fmt.Println(tab, "TAB STRING URL PATH")

	if tab[1] == "non" && tab[2] == "connecté" {
		fmt.Println("vous n'etes pas connecté")

		data := value{}
		http.Redirect(w, r, "http://localhost:8080/topic/"+tab[5], http.StatusSeeOther)
		topic.Execute(w, data)

	} else {
		dbHost := "localhost"
		dbPort := "3306"
		dbUser := "root"
		dbPass := ""
		dbName := "test"
		dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		db, err := sql.Open("mysql", dbURI)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("database is reachable")
		select2, err := db.Query("SELECT contenu,dateheure FROM `message` WHERE id = ?", tab[0])

		if err != nil {
			panic(err)
		}

		var contenumsg string
		var dateheure string
		select2.Next()

		err1 := select2.Scan(&contenumsg, &dateheure) //&id_utilisateurs, &id_messages

		if err1 != nil {

			panic(err1.Error())
		}

		defer select2.Close()

		if strings.Contains(contenumsg, "Ø") { //si le message est une réponse on enlève les séparateurs Ø
			var temp string
			var tab []string
			for _, part := range contenumsg {
				if string(part) != "Ø" {
					temp = temp + string(part)

				}
				if string(part) == "Ø" {

					fmt.Println(temp)
					tab = append(tab, string(temp))
					temp = ""
				}

			}
			contenumsg = tab[1] // le contenu du message  épuré des séparateurs
		}

		data := ""
		//génération du la div qui permet à l'utilisateur de visualiser le message  de la personne à qui il répond
		data = data + "<div class='Parent'><h1>Réponse à  " + tab[4] + " " + tab[3] + "</h1><a href ='http://localhost:8080/topic/" + tab[5] + "' style='color:none;text-decoration:none;color:white;'><div class='contenuTopic' style='margin:0rem -1rem;'><div class='DatesHeuresTopics'><p class='Heures'> Posté le  " + dateheure + "&nbsp;</p><p class='Auteur'> " + tab[4] + " " + tab[3] + "  a dit :&nbsp;</p></div><div class='ParentTitreTopics'><p class='MessagesTopics'>" + contenumsg + "</p></div></div></a><div class='ReponseMessage'><label for='mess'>Votre message<form> <button class='EnvoieNouveauxMessages' style='margin :0rem 19rem;' type='submit'>Envoyer message</button></label><div class='replyareatext'><textarea name='reply' class='textreply' >entrez votre réponse </textarea></div></form></div></div>"
		var reply = r.FormValue("reply")
		if reply == "" { // si le message est vide
			fmt.Println("le message est vide ")

			topic.Execute(w, data)
			w.Write([]byte(data))
		} else {
			var contenu string
			contenu = contenumsg + "Ø" + reply + "Ø" + tab[4] + "Ø" + tab[3] + "Ø" // on entoure chaque information message d'origine,réponse,auteur d'un séparateur
			select3, err := db.Prepare("INSERT INTO  message  (contenu,dateheure,id_topic,id_utilisateurs) VALUES(?,?,?,?);")

			if err != nil {
				panic(err)
			}

			defer select2.Close()
			var time = time.Now().Format("2006-01-02 15:04")
			_, replyto := select3.Exec(""+contenu+"", ""+time+"", ""+tab[5]+"", ""+strconv.Itoa(id_compte)+"") // on ajoute la réponse à la DB
			if replyto != nil {
				panic(replyto)
			}

			select5, err := db.Query("SELECT id FROM `message`  WHERE contenu = ?", contenu) // on récupère l'id pour la réponse

			if err != nil {
				panic(err)
			}
			var id_msg string

			select5.Next()

			err4 := select5.Scan(&id_msg)

			if err4 != nil {

				panic(err4.Error())
			}

			defer select3.Close()

			http.Redirect(w, r, "http://localhost:8080/topic/"+tab[5]+"#"+id_msg, http.StatusSeeOther) //on le redirige vers le message

			topic.Execute(w, data)
			w.Write([]byte(data))
		}
	}
}

func HttpLikeHandler(w http.ResponseWriter, r *http.Request) { // fonction de la page qui gère les likes

	fmt.Println("URLPATH - LIKE", r.URL.Path[:], len(r.URL.Path[:]))
	var cheminentier = r.URL.Path[:]
	fmt.Println(cheminentier)

	var temp string
	var tab []string

	for _, part := range cheminentier {
		if string(part) != "-" {
			temp = temp + string(part)

		}
		if temp == "/like/" { // on ignore le début de la requete
			temp = ""
		}
		if string(part) == "-" { // meme fonctionnement que sur les autres pages

			fmt.Println(temp)
			tab = append(tab, string(temp))
			temp = ""
		}

	}

	var name = tab[2]

	fmt.Println("TAB ENTIER-LIKE", tab)

	fmt.Println("ID DU COMPTE")

	if tab[1] == "non" && tab[2] == "connecté" || id_compte == 0 { // si l'utilisateur n'est pas connecté
		fmt.Println("vous n'etes pas connecté (like)")

		data := value{}
		http.Redirect(w, r, "http://localhost:8080/topic/"+tab[3]+"#"+tab[0], http.StatusSeeOther) // on le redirige vers la page d'acceuil
		topic.Execute(w, data)

	} else {

		dbHost := "localhost"
		dbPort := "3306"
		dbUser := "root"
		dbPass := ""
		dbName := "test" //nom_bdd

		// Chaîne de connexion
		dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		// Connexion à la base de données
		db, err := sql.Open("mysql", dbURI)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("NAME LIKER", name)
		fmt.Println("database is reachable", db, dbURI)

		select1, err0 := db.Query("SELECT id FROM `utilisateurs`  WHERE nom = ?", name) //on vérifie  avec ces deux requetes si l'utilisateur n'a pas déja liké
		if err0 != nil {
			panic(err0)
		}
		var idliker int // id du l'utilisateur qui a liké

		select1.Next()

		err3 := select1.Scan(&idliker)

		if err3 != nil {

			panic(err3.Error())
		}
		defer select1.Close()

		select3, err1 := db.Query("SELECT id_utilisateurs FROM `appreciation_messages` WHERE id_messages = ?", tab[0])

		if err1 != nil {
			panic(err1)
		}
		var isliked bool
		var idtocompare sql.NullInt64

		for select3.Next() {

			err4 := select3.Scan(&idtocompare)
			fmt.Println(int(idtocompare.Int64), idliker, "IDLIKER ET IDTOCOMPARE")
			if idliker == int(idtocompare.Int64) { // on vérifie si  l'utilisateur n'a pas déja liké
				isliked = true

			}

			if err4 != nil {

				panic(err3.Error())
			}
		}
		defer select3.Close()

		fmt.Println("ID USER WHO LIKED  - ID TO COMPARE", idliker, idtocompare)

		if isliked { // on redirige l'utilisateur vers le message du like  si il a déja liké
			fmt.Println("vous avez déja liké")

			data := value{}
			http.Redirect(w, r, "http://localhost:8080/topic/"+tab[3]+"#"+tab[0], http.StatusSeeOther)
			topic.Execute(w, data)

		} else {

			select1, err := db.Prepare("INSERT INTO appreciation_messages  (id_utilisateurs,id_messages,etat_like) VALUES(?,?,?);")

			if err != nil {
				panic(err)
			}

			_, registernewuser := select1.Exec(""+strconv.Itoa(idliker)+"", ""+tab[0]+"", 1) // on ajout un like au maessage dans la BD
			if registernewuser != nil {
				panic(registernewuser)
			}

			data := value{}
			http.Redirect(w, r, "http://localhost:8080/topic/"+tab[3]+"#"+tab[0], http.StatusSeeOther) //redirection
			topic.Execute(w, data)

		}

	}
}

func HttpDislikeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("URLPATH - DISLIKE", r.URL.Path[:], len(r.URL.Path[:]))
	var cheminentier = r.URL.Path[:] // réperation chemin url
	fmt.Println(cheminentier)

	var temp string
	var tab []string

	for _, part := range cheminentier { // meme utilité quand dans les autres fonctions
		if string(part) != "-" {
			temp = temp + string(part)

		}
		if temp == "/dislike/" {
			temp = ""
		}
		if string(part) == "-" {

			fmt.Println(temp)
			tab = append(tab, string(temp))
			temp = ""
		}

	}

	var name = tab[2]

	fmt.Println("TAB ENTIER-LIKE", tab)

	fmt.Println("ID DU COMPTE")

	if tab[1] == "non" && tab[2] == "connecté" || id_compte == 0 { // si l'utilisateur n'est pas connecté
		fmt.Println("vous n'etes pas connecté (dislike)")

		data := value{}
		http.Redirect(w, r, "http://localhost:8080/topic/"+tab[3]+"#"+tab[0], http.StatusSeeOther) //redirection
		topic.Execute(w, data)

	} else {

		dbHost := "localhost"
		dbPort := "3306"
		dbUser := "root"
		dbPass := ""
		dbName := "test"

		dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		db, err := sql.Open("mysql", dbURI)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("NAME DIS-LIKER", name)
		fmt.Println("database is reachable", db, dbURI)

		select1, err0 := db.Query("SELECT id FROM `utilisateurs`  WHERE nom = ?", name) //on vérifie par ces deux requetes pour  le  vérifier si l'utilisateur à deja disliké

		if err0 != nil {
			panic(err0)
		}
		var iddisliker int

		select1.Next()

		err3 := select1.Scan(&iddisliker)

		if err3 != nil {

			panic(err3.Error())
		}
		select1.Close()

		select3, err1 := db.Query("SELECT id_utilisateurs FROM `appreciation_messages` WHERE id_messages = ?", tab[0])

		if err1 != nil {
			panic(err1)
		}
		var isdisliked bool
		var idtocompare sql.NullInt64

		for select3.Next() {

			err4 := select3.Scan(&idtocompare)
			fmt.Println(int(idtocompare.Int64), iddisliker, "IDLIKER ET IDTOCOMPARE")
			if iddisliker == int(idtocompare.Int64) { // si il y a déja  une occurence de l'id utilisateur pour le message dans la table appreciation_message  (si l'utilisateur a déja disliké)
				isdisliked = true

			}
			fmt.Println(idtocompare)

			if err4 != nil {

				panic(err3.Error())
			}
		}
		defer select3.Close()

		fmt.Println("ID USER WHO LIKED  - ID TO COMPARE", iddisliker, idtocompare)

		if isdisliked {
			fmt.Println("vous avez déja disliké")

			data := value{}
			http.Redirect(w, r, "http://localhost:8080/topic/"+tab[3]+"#"+tab[0], http.StatusSeeOther) //redirection
			topic.Execute(w, data)

		} else {

			select1, err := db.Prepare("INSERT INTO appreciation_messages  (id_utilisateurs,id_messages,etat_like) VALUES(?,?,?);")

			if err != nil {
				panic(err)
			}

			_, registernewuser := select1.Exec(""+strconv.Itoa(iddisliker)+"", ""+tab[0]+"", 0) //ajout du dislike dans la bd
			if registernewuser != nil {
				panic(registernewuser)
			}

			data := value{}
			http.Redirect(w, r, "http://localhost:8080/topic/"+tab[3]+"#"+tab[0], http.StatusSeeOther) //redirection
			topic.Execute(w, data)

		}

	}

	data := value{}
	topic.Execute(w, data)

}

func HttpCGUHandler(w http.ResponseWriter, r *http.Request) { //gestion de la page cgu

	data := value{}
	cgu.Execute(w, data)

}

func HttpPolicyHandler(w http.ResponseWriter, r *http.Request) { // gestion de la page policy

	data := value{}
	policy.Execute(w, data)

}

func HttpAboutHandler(w http.ResponseWriter, r *http.Request) { // gestion de la page about

	data := value{}
	about.Execute(w, data)

}

func HttpContactUsHandler(w http.ResponseWriter, r *http.Request) { // gestion de la page handler

	data := value{}
	contactus.Execute(w, data)

}
