// +build bench

package hw10_program_optimization //nolint:golint,stylecheck

import (
	"archive/zip"
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	mb          uint64 = 1 << 20
	memoryLimit uint64 = 30 * mb

	timeLimit = 300 * time.Millisecond
)

func benchmarkGetDomainStatShort(b *testing.B) {
	data := `{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
	{"Id":2,"Name":"Justin Oliver Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"oPerez","Email":"MelissaGutierrez@Twinte.biz","Phone":"106-05-18","Password":"f00GKr9i","Address":"Oak Valley Lane 19"}
	{"Id":3,"Name":"Brian Olson","Username":"non_quia_id","Email":"FrancesEllis@Quinu.edu","Phone":"237-75-34","Password":"cmEPhX8","Address":"Butterfield Junction 74"}
	{"Id":4,"Name":"Jesse Vasquez Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"qRichardson","Email":"mLynch@Dabtype.name","Phone":"9-373-949-64-00","Password":"SiZLeNSGn","Address":"Fulton Hill 80"}
	{"Id":5,"Name":"Clarence Olson","Username":"RachelAdams","Email":"RoseSmith@Browsecat.com","Phone":"988-48-97","Password":"71kuz3gA5w","Address":"Monterey Park 39"}
	{"Id":6,"Name":"Gregory Reid","Username":"tButler","Email":"5Moore@Teklist.net","Phone":"520-04-16","Password":"r639qLNu","Address":"Sunfield Park 20"}
	{"Id":7,"Name":"Janice Rose","Username":"KeithHart","Email":"nulla@Linktype.com","Phone":"146-91-01","Password":"acSBF5","Address":"Russell Trail 61"}
	{"Id":8,"Name":"Jacqueline Young","Username":"CraigKnight","Email":"kCunningham@Skiptube.gov","Phone":"6-954-746-32-77","Password":"rHBCvD5JpLGs","Address":"4th Pass 91"}
	{"Id":9,"Name":"Steve Burns","Username":"bRoberts","Email":"perferendis@Skippad.name","Phone":"246-85-85","Password":"68xyVtL1AaO6","Address":"Jenifer Circle 24"}
	{"Id":10,"Name":"Paula Gonzales","Username":"4Ramirez","Email":"BrianBradley@Zoomcast.info","Phone":"363-62-16","Password":"atpnGIr","Address":"Barnett Park 43"}
	{"Id":11,"Name":"Janet Clark I II III IV V MD DDS PhD DVM","Username":"8Perry","Email":"SaraHawkins@Eazzy.edu","Phone":"2-306-840-60-85","Password":"h7atkNURvN","Address":"Bunting Lane 40"}
	{"Id":12,"Name":"Angela Nichols","Username":"quia","Email":"qui@Topiczoom.org","Phone":"9-591-154-17-64","Password":"lorrSrfSaqxk","Address":"Twin Pines Alley 66"}
	{"Id":13,"Name":"Roger Gilbert","Username":"doloremque","Email":"et@Riffwire.info","Phone":"0-710-227-54-55","Password":"1lW0kUPXxj","Address":"Mifflin Lane 55"}
	{"Id":14,"Name":"Donna Vasquez","Username":"aspernatur","Email":"GeraldHughes@Rhycero.name","Phone":"737-84-46","Password":"9d3cys3VNc","Address":"Lakewood Center 85"}
	{"Id":15,"Name":"Todd Payne","Username":"dWilliams","Email":"3Kennedy@Plajo.gov","Phone":"0-915-255-08-38","Password":"DjkShy9NQ1R","Address":"Gulseth Parkway 72"}
	{"Id":16,"Name":"Frances Olson","Username":"9Meyer","Email":"et_et@Twitterlist.biz","Phone":"0-402-112-61-97","Password":"ioOjbYTZd","Address":"Kipling Trail 73"}
	{"Id":17,"Name":"Diana Palmer","Username":"quia_omnis_temporibus","Email":"AliceAustin@Realblab.biz","Phone":"288-12-53","Password":"Yyh5pLYvO7K","Address":"Butterfield Hill 19"}
	{"Id":18,"Name":"Dorothy Bradley","Username":"SharonGarza","Email":"VirginiaPrice@Photospace.info","Phone":"833-89-59","Password":"TILpRLI","Address":"Judy Center 97"}
	{"Id":19,"Name":"Carl Crawford","Username":"aut","Email":"dWilliams@Cogilith.info","Phone":"3-370-475-78-17","Password":"IX1rDUxz1","Address":"6th Terrace 36"}
	{"Id":20,"Name":"Denise Roberts","Username":"repudiandae","Email":"quia_iusto_laboriosam@Roomm.org","Phone":"7-043-762-06-95","Password":"Ej6DnzIO5","Address":"Burrows Trail 17"}
	{"Id":21,"Name":"Daniel Jenkins","Username":"PhilipFrazier","Email":"RichardPeterson@Gigashots.org","Phone":"981-77-46","Password":"JfaOx58","Address":"Morningstar Alley 24"}
	{"Id":22,"Name":"Timothy Gilbert","Username":"praesentium_ut_et","Email":"et@Chatterpoint.biz","Phone":"826-94-61","Password":"VrUooZL8F8cN","Address":"Orin Plaza 40"}
	{"Id":23,"Name":"Kimberly Jackson","Username":"MatthewMorales","Email":"2Snyder@Pixonyx.org","Phone":"7-452-214-79-06","Password":"wgMiOgV4","Address":"Hintze Alley 48"}
	{"Id":24,"Name":"Alan Kelley","Username":"PatriciaGilbert","Email":"totam_excepturi_dolore@Oozz.biz","Phone":"9-735-960-25-03","Password":"BiDTcevghYo","Address":"Fulton Terrace 76"}
	{"Id":25,"Name":"Joan Mills","Username":"in_similique_pariatur","Email":"1Webb@Rhynyx.mil","Phone":"868-93-23","Password":"uC2OdY","Address":"Bartillon Lane 57"}
	{"Id":26,"Name":"Thomas Carpenter","Username":"qui_delectus_optio","Email":"jDay@Photobug.com","Phone":"4-598-198-74-05","Password":"BvKJQRFo","Address":"Northland Circle 95"}
	{"Id":27,"Name":"Willie Matthews","Username":"mHenderson","Email":"autem_et_rerum@Demimbu.org","Phone":"392-26-45","Password":"0NNbpcv0oxG9","Address":"Lakeland Plaza 22"}
	{"Id":28,"Name":"Rachel Andrews","Username":"yGreen","Email":"rerum@Flashpoint.mil","Phone":"7-672-375-42-30","Password":"LfEmdgN2","Address":"Melvin Hill 50"}
	{"Id":29,"Name":"Laura Allen","Username":"HenryAustin","Email":"vLewis@Jabbersphere.info","Phone":"1-047-724-35-40","Password":"fyzHcHatSz4U","Address":"Mendota Parkway 20"}
	{"Id":30,"Name":"Lillian Cruz","Username":"PatrickCampbell","Email":"CynthiaWashington@Pixoboo.com","Phone":"992-74-36","Password":"0pJyIl1C","Address":"Clove Hill 64"}
	{"Id":31,"Name":"Mr. Dr. Terry Walker","Username":"neque","Email":"RoseMason@Blognation.name","Phone":"8-663-274-08-44","Password":"7YtsP0skeJ","Address":"Harbort Circle 60"}
	{"Id":32,"Name":"Russell Stephens","Username":"DanielHall","Email":"voluptas_fugiat_perferendis@Devpoint.org","Phone":"657-59-67","Password":"FrLNI59Bl4ug","Address":"Mendota Avenue 93"}
	{"Id":33,"Name":"Kelly Hudson","Username":"xGomez","Email":"et@Quire.edu","Phone":"972-75-35","Password":"SscH9ZH1","Address":"Green Plaza 92"}
	{"Id":34,"Name":"Mrs. Ms. Miss Nicole Mills","Username":"5Hanson","Email":"qBailey@Vitz.mil","Phone":"601-52-39","Password":"ztehlTSg","Address":"Walton Park 30"}
	{"Id":35,"Name":"Ryan Gilbert","Username":"quibusdam_rem","Email":"TheresaGomez@Brainlounge.biz","Phone":"1-454-103-67-34","Password":"Pf775knRvuIt","Address":"Raven Hill 0"}
	{"Id":36,"Name":"Catherine Morgan","Username":"KevinLewis","Email":"zHudson@Dabjam.gov","Phone":"228-47-80","Password":"JOecQz","Address":"Del Mar Street 11"}
	{"Id":37,"Name":"Gregory Baker","Username":"ea","Email":"quis_quo@Brainlounge.name","Phone":"797-38-46","Password":"bQZK8RR","Address":"Messerschmidt Hill 16"}
	{"Id":38,"Name":"Richard Cooper","Username":"nihil","Email":"NormaLawson@Izio.gov","Phone":"5-982-792-78-04","Password":"3TdZMYabH1lu","Address":"North Trail 49"}
	{"Id":39,"Name":"Peter Nguyen","Username":"bOwens","Email":"mMiller@Quaxo.mil","Phone":"469-60-41","Password":"Kni3LpwID","Address":"High Crossing Center 57"}
	{"Id":40,"Name":"Joan Gonzalez","Username":"MarthaFox","Email":"qui_maxime@Livetube.info","Phone":"858-94-05","Password":"QUzPHIOUfw04","Address":"Muir Parkway 67"}
	{"Id":41,"Name":"Lillian Gonzales","Username":"JaniceKennedy","Email":"JudithBell@Feedfire.edu","Phone":"085-84-82","Password":"m0Gr6PFprt","Address":"Arapahoe Circle 57"}
	{"Id":42,"Name":"Dennis Ryan","Username":"EdwardBurton","Email":"IreneFerguson@Omba.gov","Phone":"0-755-554-72-24","Password":"gyEbmzyZG9p","Address":"Johnson Way 12"}
	{"Id":43,"Name":"Jerry Hall","Username":"cupiditate","Email":"nesciunt_ipsa@Feedbug.biz","Phone":"0-523-541-57-20","Password":"MA2DQt","Address":"Arapahoe Park 14"}
	{"Id":44,"Name":"Diana Mitchell","Username":"FrancesElliott","Email":"ipsam_itaque_possimus@Browsedrive.org","Phone":"0-480-969-03-61","Password":"fdTfF03RMKyb","Address":"Dunning Point 22"}
	{"Id":45,"Name":"Debra Stone","Username":"MarthaRice","Email":"kMorgan@Meetz.net","Phone":"589-04-83","Password":"JXXrjK1kW","Address":"Chive Circle 21"}
	{"Id":46,"Name":"Jane White","Username":"dolor_commodi","Email":"FrancesMyers@Jaxworks.name","Phone":"298-72-50","Password":"SOGYpYXGzX","Address":"Farwell Circle 69"}
	{"Id":47,"Name":"Nicholas Carter","Username":"dolorem_neque","Email":"voluptas_expedita@Fivechat.com","Phone":"4-395-019-81-56","Password":"evEZ1tIyE3B8","Address":"Doe Crossing Pass 94"}
	{"Id":48,"Name":"Charles Mcdonald","Username":"LawrenceChapman","Email":"GregoryVasquez@Dynabox.edu","Phone":"4-433-618-25-42","Password":"rdjpUdsP","Address":"Ridgeway Point 64"}
	{"Id":49,"Name":"Ryan Fox","Username":"pariatur_rerum","Email":"DebraAustin@Camido.net","Phone":"538-49-70","Password":"XOGImn","Address":"Corben Parkway 27"}
	{"Id":50,"Name":"Ashley Powell","Username":"oWheeler","Email":"7Alexander@Brainverse.edu","Phone":"3-596-905-78-45","Password":"VNiaUZSAJ","Address":"Northport Junction 61"}
	{"Id":51,"Name":"Walter Webb","Username":"magnam_illo","Email":"sBradley@Eamia.org","Phone":"460-16-18","Password":"5V8jZ86Lr","Address":"Utah Avenue 23"}
	{"Id":52,"Name":"Jeffrey Palmer","Username":"quaerat_aut","Email":"9Stone@Yodoo.info","Phone":"537-48-95","Password":"LgV8PiW","Address":"Arrowood Lane 38"}
	{"Id":53,"Name":"Chris Daniels","Username":"JuanNelson","Email":"4Snyder@Eimbee.net","Phone":"847-32-47","Password":"DvhtmWlcCtgV","Address":"Grover Circle 93"}
	{"Id":54,"Name":"Eugene Ross Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"rerum_et","Email":"incidunt_ut@Photobug.gov","Phone":"362-96-90","Password":"bBsAShPK","Address":"Mcbride Pass 26"}
	{"Id":55,"Name":"Carol Price","Username":"et","Email":"DianeGarcia@Ainyx.org","Phone":"374-42-57","Password":"H4Ee7N","Address":"Ohio Terrace 11"}
	{"Id":56,"Name":"Angela Kennedy","Username":"ad","Email":"hWalker@Muxo.biz","Phone":"035-83-54","Password":"z3O2FUI","Address":"Delladonna Trail 25"}
	{"Id":57,"Name":"Lori Mcdonald","Username":"0Garcia","Email":"gMarshall@Riffwire.net","Phone":"6-277-143-52-44","Password":"H2RaBsW6t","Address":"Banding Point 80"}
	{"Id":58,"Name":"Joan Cook","Username":"AlbertCrawford","Email":"uBurton@Kare.name","Phone":"2-131-557-17-74","Password":"BAXIJew6Nl","Address":"Granby Place 83"}
	{"Id":59,"Name":"Stephen Harvey","Username":"magnam_quia_omnis","Email":"possimus@Meetz.mil","Phone":"550-84-86","Password":"3v1KM4mOkzk7","Address":"Pepper Wood Court 37"}
	{"Id":60,"Name":"Carl Hart","Username":"MarieRuiz","Email":"gMatthews@Eazzy.gov","Phone":"7-326-274-13-70","Password":"tW7ABC0kKKG","Address":"Jackson Place 52"}
	{"Id":61,"Name":"Katherine Gonzales","Username":"aut","Email":"aNelson@Shufflester.name","Phone":"3-425-855-49-63","Password":"eBBb8yhG","Address":"Luster Circle 96"}
	{"Id":62,"Name":"Edward Ryan","Username":"ElizabethMartin","Email":"JenniferReid@Aimbu.edu","Phone":"531-39-23","Password":"3DHzNQ","Address":"Buhler Court 75"}
	{"Id":63,"Name":"Marie Castillo","Username":"omnis","Email":"ElizabethHarper@Quire.net","Phone":"7-109-693-15-09","Password":"bOm4FJ","Address":"Ryan Park 77"}
	{"Id":64,"Name":"Stephen Walker","Username":"rJohnston","Email":"1Gonzales@Npath.com","Phone":"3-454-523-79-45","Password":"yhmtsEbHHD","Address":"Cherokee Trail 66"}
	{"Id":65,"Name":"Samuel Morales","Username":"SandraPerkins","Email":"JanetPerez@Yabox.net","Phone":"723-64-81","Password":"JyXiOFqbuRD","Address":"Union Park 11"}
	{"Id":66,"Name":"Anne Holmes","Username":"sed","Email":"aut_accusamus_voluptates@Linkbuzz.net","Phone":"309-17-70","Password":"Nh2Dd2OP","Address":"Caliangt Parkway 49"}
	{"Id":67,"Name":"Matthew Dean","Username":"ut_fugiat","Email":"JenniferGarcia@Mynte.name","Phone":"482-73-56","Password":"vgZAcI","Address":"Arkansas Park 6"}
	{"Id":68,"Name":"Brenda Robertson","Username":"gKing","Email":"kMyers@Photobug.net","Phone":"7-231-888-19-58","Password":"c92iMhkh87","Address":"Buena Vista Parkway 65"}
	{"Id":69,"Name":"Randy King","Username":"9Clark","Email":"illum@Skilith.org","Phone":"493-65-62","Password":"8ViH2yQJF8O","Address":"Shopko Alley 10"}
	{"Id":70,"Name":"Sharon Hall","Username":"JenniferDuncan","Email":"perferendis_sunt_nihil@Lazzy.mil","Phone":"771-65-53","Password":"W85lRiWkKk","Address":"Fieldstone Pass 16"}
	{"Id":71,"Name":"Helen White","Username":"dolores_laboriosam","Email":"voluptatem_aspernatur@Ooba.gov","Phone":"848-95-44","Password":"5MX4UKmMTPO","Address":"Reinke Drive 10"}
	{"Id":72,"Name":"Lori Jones","Username":"soluta","Email":"ThomasJacobs@Trilith.mil","Phone":"597-02-42","Password":"JwOZxx6fa3","Address":"Dayton Pass 38"}
	{"Id":73,"Name":"Sandra Fuller","Username":"accusamus_dolorem_corporis","Email":"wSchmidt@Izio.net","Phone":"9-843-985-75-86","Password":"1rLKAygJjXHy","Address":"Clove Hill 16"}
	{"Id":74,"Name":"Mr. Dr. Matthew Duncan","Username":"eveniet_maiores","Email":"repudiandae_delectus@Voomm.net","Phone":"053-85-72","Password":"zvQn1Q1V","Address":"Acker Drive 44"}
	{"Id":75,"Name":"Philip Fernandez","Username":"non_suscipit_quia","Email":"6Schmidt@Zoozzy.biz","Phone":"431-93-89","Password":"CMpTvY","Address":"Westend Pass 75"}
	{"Id":76,"Name":"Fred Riley","Username":"quia","Email":"AnnieHart@Yambee.mil","Phone":"5-628-099-02-96","Password":"3WxoUba2dRC","Address":"Becker Plaza 53"}
	{"Id":77,"Name":"Katherine Lane","Username":"JimmyHill","Email":"velit_consequatur@Brightbean.mil","Phone":"5-646-145-01-98","Password":"6WuFTEfNnc","Address":"Northport Drive 65"}
	{"Id":78,"Name":"Beverly Frazier","Username":"WayneFerguson","Email":"earum_aspernatur@Meejo.info","Phone":"0-773-641-87-60","Password":"2YnfLqAjK","Address":"Thackeray Parkway 45"}
	{"Id":79,"Name":"Jeremy Davis","Username":"numquam_laudantium","Email":"at_labore@Skimia.info","Phone":"6-168-796-00-99","Password":"AOTBvj","Address":"Jenifer Drive 32"}
	{"Id":80,"Name":"Jack Greene","Username":"et_officia","Email":"zFuller@Oodoo.mil","Phone":"0-731-528-34-66","Password":"TLh50d","Address":"Moland Alley 10"}
	{"Id":81,"Name":"Kimberly Reynolds","Username":"uHunter","Email":"1Payne@Brightdog.mil","Phone":"109-83-48","Password":"D7KdJjQg","Address":"Darwin Trail 67"}
	{"Id":82,"Name":"Lois Nelson I II III IV V MD DDS PhD DVM","Username":"LarryBrooks","Email":"6Burton@Wikizz.biz","Phone":"232-80-54","Password":"HKCuSC4Yqgs","Address":"Pankratz Circle 89"}
	{"Id":83,"Name":"Julia Fisher","Username":"2Martin","Email":"NicholasSpencer@Skilith.mil","Phone":"979-22-20","Password":"ysGxngNj3lQ","Address":"Susan Drive 9"}
	{"Id":84,"Name":"Brenda Owens","Username":"qGraham","Email":"totam_incidunt_mollitia@Zoovu.name","Phone":"508-32-18","Password":"owisqpoBtiU","Address":"Mayer Place 55"}
	{"Id":85,"Name":"Howard Chavez","Username":"AnthonyHoward","Email":"EdwardRamirez@Babblestorm.com","Phone":"4-844-996-28-66","Password":"4PA2MPde8","Address":"Marquette Parkway 60"}
	{"Id":86,"Name":"Kimberly Cooper I II III IV V MD DDS PhD DVM","Username":"quis","Email":"HaroldHamilton@Zooveo.name","Phone":"3-547-932-04-86","Password":"xbazaMah9hjB","Address":"Corry Place 31"}
	{"Id":87,"Name":"Stephanie Watson","Username":"zMurray","Email":"9Peterson@Linkbridge.biz","Phone":"7-670-678-15-47","Password":"HVfJrL","Address":"3rd Point 85"}
	{"Id":88,"Name":"Matthew Rose","Username":"RobinMason","Email":"JeffreyLong@Realfire.biz","Phone":"0-024-932-96-29","Password":"SuDz0x","Address":"Shasta Drive 30"}
	{"Id":89,"Name":"Sean Davis","Username":"fWeaver","Email":"PaulaDean@Skinte.info","Phone":"693-72-74","Password":"5k1gligU","Address":"Troy Lane 19"}
	{"Id":90,"Name":"Christine Reynolds I II III IV V MD DDS PhD DVM","Username":"voluptatem_eum","Email":"JamesParker@Eare.name","Phone":"722-34-59","Password":"TCfOIBlLkmo9","Address":"Maywood Parkway 8"}
	{"Id":91,"Name":"Tammy Willis","Username":"eos_non_facere","Email":"dolorum@Quatz.name","Phone":"4-301-664-88-76","Password":"ynPeOKdFNHY","Address":"Buhler Alley 63"}
	{"Id":92,"Name":"Walter Russell","Username":"nulla_dolores","Email":"est_maiores@Kwilith.name","Phone":"5-245-087-14-38","Password":"zJX5yBS","Address":"Fordem Pass 87"}
	{"Id":93,"Name":"Marie Smith","Username":"AdamWelch","Email":"GaryHayes@Agimba.mil","Phone":"1-143-372-01-13","Password":"mfyyO4PK6","Address":"Harbort Alley 48"}
	{"Id":94,"Name":"Mr. Dr. Harold Harrison","Username":"JeanGreen","Email":"JoyceBarnes@Voonder.biz","Phone":"697-33-84","Password":"iVKqnVm","Address":"Lillian Plaza 65"}
	{"Id":95,"Name":"Donna Frazier","Username":"architecto","Email":"JaneMartinez@Dynabox.name","Phone":"554-87-27","Password":"sa3uWfh8j","Address":"Mockingbird Court 27"}
	{"Id":96,"Name":"Michael Castillo","Username":"iAdams","Email":"pWalker@Oba.gov","Phone":"7-577-071-18-31","Password":"d48oKNW6","Address":"Hoffman Park 61"}
	{"Id":97,"Name":"Frank Roberts","Username":"soluta_eveniet","Email":"RachelJones@Shuffletag.name","Phone":"0-798-591-73-07","Password":"pkqKTTMpD","Address":"Westridge Court 81"}
	{"Id":98,"Name":"Jimmy Carpenter Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"tempore","Email":"cumque@Dazzlesphere.net","Phone":"235-87-91","Password":"NpTY6oS","Address":"Old Shore Junction 83"}
	{"Id":99,"Name":"Mary Peterson","Username":"tempore","Email":"laboriosam@Zoonoodle.info","Phone":"915-06-45","Password":"58mgvl","Address":"Jana Circle 59"}
	{"Id":100,"Name":"Brandon Fowler","Username":"AlbertBarnes","Email":"JeffreyGibson@Jatri.info","Phone":"6-833-399-61-00","Password":"t4UeINnCF","Address":"Porter Avenue 16"}
	{"Id":101,"Name":"Mrs. Ms. Miss Betty George","Username":"StevenRodriguez","Email":"fJenkins@Innotype.edu","Phone":"769-26-60","Password":"Du8KzuoqxJZ5","Address":"Merchant Lane 22"}
	{"Id":102,"Name":"Richard Washington Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"veniam_in_omnis","Email":"KatherineReynolds@Oyope.info","Phone":"0-779-715-97-11","Password":"ujlfFL3du29a","Address":"Dexter Trail 94"}
	{"Id":103,"Name":"Ruby Edwards","Username":"qui_est_error","Email":"quis_voluptatibus@Divanoodle.net","Phone":"325-65-56","Password":"xMSz99FgE","Address":"Colorado Circle 15"}
	{"Id":104,"Name":"Clarence Hamilton","Username":"cum_tempora","Email":"EmilyMartin@Kwinu.com","Phone":"706-05-34","Password":"j2RZJH","Address":"Scott Terrace 1"}
	{"Id":105,"Name":"Alan Murphy","Username":"ea_molestiae_eos","Email":"FredShaw@Jabbersphere.org","Phone":"297-13-04","Password":"T3wU4wR","Address":"Ramsey Road 86"}
	{"Id":106,"Name":"Rebecca Gibson","Username":"BettyGray","Email":"aMeyer@Gabspot.gov","Phone":"3-305-940-68-22","Password":"AVVflLZ6","Address":"Knutson Circle 87"}
	{"Id":107,"Name":"Alan Freeman Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"rDay","Email":"BillyGordon@Ozu.info","Phone":"298-70-16","Password":"gW5yEhV","Address":"Colorado Parkway 49"}
	{"Id":108,"Name":"Randy Miller","Username":"LarryRivera","Email":"AnnWest@Vinder.info","Phone":"833-53-97","Password":"0mkzxu","Address":"Sycamore Avenue 13"}
	{"Id":109,"Name":"Joyce Griffin","Username":"ipsam_dolore_eum","Email":"wWashington@Tavu.gov","Phone":"5-954-907-16-81","Password":"yKmg312X","Address":"Rutledge Alley 44"}
	{"Id":110,"Name":"Norma Garza","Username":"rJames","Email":"JustinRodriguez@Quinu.biz","Phone":"7-593-012-61-06","Password":"D8cYVpzP","Address":"Del Mar Way 89"}
	{"Id":111,"Name":"Billy Welch","Username":"BeverlyKim","Email":"sit_vel_est@Layo.mil","Phone":"5-264-025-13-48","Password":"yn3c38vFHZjY","Address":"Debs Court 22"}
	{"Id":112,"Name":"Irene Bell","Username":"sHughes","Email":"CherylRivera@Aimbo.name","Phone":"4-250-495-46-53","Password":"jDZDtY2Kxx","Address":"Talisman Avenue 61"}
	{"Id":113,"Name":"Martha Bryant","Username":"repellendus","Email":"ChristinaGarrett@Agivu.gov","Phone":"8-221-786-33-13","Password":"wqlsVKbN4c","Address":"Walton Pass 4"}
	{"Id":114,"Name":"Rachel Washington","Username":"9Ryan","Email":"libero_repellat_placeat@Abatz.biz","Phone":"4-093-712-41-06","Password":"qzjhre78","Address":"Dwight Plaza 56"}
	{"Id":115,"Name":"Joseph West","Username":"BrendaShaw","Email":"omnis@Jetpulse.com","Phone":"251-46-17","Password":"hhJDNbrepa","Address":"Fulton Parkway 8"}
	{"Id":116,"Name":"Thomas Wallace","Username":"3Day","Email":"vel_culpa@Fliptune.name","Phone":"9-962-905-15-86","Password":"Gfvss1S7","Address":"Annamark Court 37"}
	{"Id":117,"Name":"Jerry Washington","Username":"0Ortiz","Email":"ipsum@Demivee.edu","Phone":"4-717-423-11-34","Password":"bN8DxH37NMF","Address":"Bellgrove Place 42"}
	{"Id":118,"Name":"Chris Carr","Username":"fDunn","Email":"KathrynGriffin@Cogibox.com","Phone":"2-448-562-94-40","Password":"d1eB0CPKsgj3","Address":"Atwood Place 37"}
	{"Id":119,"Name":"Paul Spencer","Username":"JeanSnyder","Email":"zAlexander@Jabbercube.com","Phone":"381-31-06","Password":"YAjRWZdE163","Address":"Dorton Way 60"}
	{"Id":120,"Name":"Louis Garrett","Username":"perferendis","Email":"yAlexander@Zoonoodle.biz","Phone":"2-785-005-19-45","Password":"SciXPpLO","Address":"Mendota Parkway 35"}
	{"Id":121,"Name":"Theresa Cook","Username":"StephanieGibson","Email":"pWood@Skiptube.biz","Phone":"4-090-034-68-59","Password":"EMP1CEdrUc7","Address":"Johnson Crossing 68"}
	{"Id":122,"Name":"Mr. Dr. Steven Lynch","Username":"illo_dicta_asperiores","Email":"TammyWalker@Thoughtstorm.info","Phone":"8-223-072-56-07","Password":"htNr54","Address":"Granby Crossing 24"}
	{"Id":123,"Name":"Emily Shaw I II III IV V MD DDS PhD DVM","Username":"IreneAlexander","Email":"GregoryJackson@Leexo.name","Phone":"029-09-96","Password":"m6vYDimE","Address":"Eggendart Drive 18"}
	{"Id":124,"Name":"Anne Vasquez","Username":"StevenHudson","Email":"DianeDuncan@Photobug.name","Phone":"013-22-02","Password":"z50gMW5","Address":"Emmet Crossing 24"}
	{"Id":125,"Name":"Martha Ramos","Username":"iFord","Email":"sequi_neque@Flashpoint.mil","Phone":"683-88-27","Password":"P1A2o7m","Address":"Hanover Crossing 88"}
	{"Id":126,"Name":"Brandon White","Username":"MarieWallace","Email":"officiis_sapiente@Katz.mil","Phone":"6-064-222-60-03","Password":"QZKaahlH8kI","Address":"Cambridge Point 16"}
	{"Id":127,"Name":"Virginia Roberts I II III IV V MD DDS PhD DVM","Username":"magnam_sequi","Email":"AnneRichardson@Ntags.net","Phone":"9-031-763-14-74","Password":"XNQTZnVjddN","Address":"Caliangt Junction 56"}
	{"Id":128,"Name":"Marilyn Kelley I II III IV V MD DDS PhD DVM","Username":"dolore_laborum_sit","Email":"nWood@Jetpulse.com","Phone":"670-27-04","Password":"DUpdwr5UG0r1","Address":"Randy Court 45"}
	{"Id":129,"Name":"Cheryl Parker","Username":"nihil","Email":"RichardBradley@Bubbletube.biz","Phone":"478-42-45","Password":"ilptxT","Address":"Mockingbird Pass 15"}
	{"Id":130,"Name":"Teresa Arnold","Username":"unde","Email":"RebeccaGray@Gabtune.name","Phone":"5-574-881-90-21","Password":"oR4TDYHS0cxM","Address":"Luster Terrace 84"}
	{"Id":131,"Name":"Andrea Mitchell","Username":"FredMorris","Email":"LouisMyers@Twiyo.info","Phone":"572-71-65","Password":"ifDbStywFqo","Address":"Hanover Terrace 21"}
	{"Id":132,"Name":"Carolyn Perkins","Username":"8Turner","Email":"facere_consequuntur_quas@Riffpedia.info","Phone":"8-836-926-07-92","Password":"zCM50nf4Gf","Address":"Cardinal Pass 82"}
	{"Id":133,"Name":"Stephanie Gibson","Username":"fuga_error_tempore","Email":"nobis_nam@Quimba.org","Phone":"510-76-32","Password":"NJaRLagBtY","Address":"Roth Trail 12"}
	{"Id":134,"Name":"Mr. Dr. Donald Garza","Username":"6Meyer","Email":"maiores_qui@Skidoo.com","Phone":"7-956-724-51-42","Password":"LkmY19AJ3","Address":"Derek Street 6"}
	{"Id":135,"Name":"Henry Williamson","Username":"KarenMoore","Email":"LoisMcdonald@Geba.com","Phone":"5-819-090-23-14","Password":"sByyrXzdRxF","Address":"Buell Circle 89"}
	{"Id":136,"Name":"Julia Holmes","Username":"qBradley","Email":"iOliver@Quatz.edu","Phone":"488-15-23","Password":"i7lR0uJBaBC","Address":"Ludington Center 71"}
	{"Id":137,"Name":"Annie Hill","Username":"3Ellis","Email":"JanetAllen@Shufflebeat.gov","Phone":"8-913-575-54-68","Password":"cfg44a","Address":"Aberg Terrace 1"}
	{"Id":138,"Name":"Patricia Hunt","Username":"LillianHart","Email":"MargaretHenderson@Blogtags.biz","Phone":"767-40-85","Password":"qwfrMotndN","Address":"Grasskamp Street 72"}
	{"Id":139,"Name":"Mr. Dr. Harry Davis","Username":"5Cook","Email":"RandyLee@Skyble.name","Phone":"315-59-61","Password":"Iv4fDgOG5","Address":"Westridge Avenue 39"}
	{"Id":140,"Name":"Mrs. Ms. Miss Janet Banks","Username":"oCollins","Email":"molestiae@Jabbercube.gov","Phone":"2-289-516-00-25","Password":"DK2cSUzUyN","Address":"Grim Road 9"}
	{"Id":141,"Name":"Martha Stewart","Username":"JerryOlson","Email":"quisquam_voluptas_repellendus@Geba.info","Phone":"3-340-069-86-18","Password":"bGV08B","Address":"Del Mar Center 67"}
	{"Id":142,"Name":"Kenneth Jackson Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"aut","Email":"exercitationem@Katz.com","Phone":"708-84-40","Password":"vqtU2Q","Address":"Longview Circle 57"}
	{"Id":143,"Name":"Walter Gardner","Username":"MildredDixon","Email":"rerum@Meezzy.info","Phone":"2-578-993-72-28","Password":"kpHczZg22E","Address":"Dryden Way 63"}
	{"Id":144,"Name":"Virginia Williams","Username":"unde","Email":"jDunn@Buzzshare.net","Phone":"9-760-385-37-44","Password":"AUcDDhz2Va","Address":"Park Meadow Alley 87"}
	{"Id":145,"Name":"Diana Patterson","Username":"9Mason","Email":"nam_et@Fanoodle.com","Phone":"542-90-12","Password":"nVhEcJ","Address":"Ryan Drive 38"}
	{"Id":146,"Name":"Douglas Garrett","Username":"RobinFisher","Email":"error@Mudo.gov","Phone":"787-85-93","Password":"JbLJOxW6f","Address":"Karstens Terrace 55"}
	{"Id":147,"Name":"Earl Howell","Username":"nWebb","Email":"DennisPerry@Jaxworks.net","Phone":"001-43-71","Password":"IYtPAskC7H","Address":"Anhalt Parkway 71"}
	{"Id":148,"Name":"Jane Peters","Username":"8Rogers","Email":"aut_illum@Skipstorm.biz","Phone":"0-434-688-70-66","Password":"Cd3ionPHB0mQ","Address":"North Street 24"}
	{"Id":149,"Name":"Martin Hall","Username":"quia","Email":"LillianAndrews@Photobug.gov","Phone":"8-405-749-08-39","Password":"1Ab492WDYSj","Address":"Almo Trail 19"}
	{"Id":150,"Name":"Doris Turner","Username":"ClarenceRichards","Email":"1Hill@Demizz.net","Phone":"718-85-13","Password":"8aBHCimnL","Address":"Melvin Alley 0"}
	{"Id":151,"Name":"Lawrence Cole","Username":"AnnaRiley","Email":"qJohnson@Aibox.biz","Phone":"799-42-87","Password":"OehpQQ7NdcID","Address":"Butternut Park 53"}
	{"Id":152,"Name":"Julia Little","Username":"voluptate_autem","Email":"LarryHill@Linklinks.gov","Phone":"8-659-651-61-26","Password":"CffPovk","Address":"Coolidge Park 7"}
	{"Id":153,"Name":"Barbara Peterson I II III IV V MD DDS PhD DVM","Username":"wMorgan","Email":"fRice@Babbleblab.com","Phone":"326-27-31","Password":"n1ZYVNGOw","Address":"Dennis Drive 53"}
	{"Id":154,"Name":"Mr. Dr. Justin Shaw","Username":"JuliaMoreno","Email":"JaniceHoward@Linkbridge.info","Phone":"0-441-359-92-04","Password":"WMpHcFJXO","Address":"Walton Circle 42"}
	{"Id":155,"Name":"Julie Walker","Username":"MarkDiaz","Email":"ipsam_odio@Edgeify.mil","Phone":"9-749-783-67-33","Password":"Qvujxn3","Address":"Nobel Road 49"}
	{"Id":156,"Name":"Helen Ryan","Username":"5Hamilton","Email":"qui_aut@Oba.info","Phone":"616-12-53","Password":"nOrXCkA24nJ","Address":"Manley Trail 46"}
	{"Id":157,"Name":"Melissa Taylor","Username":"hMedina","Email":"CharlesJones@Kwideo.gov","Phone":"5-169-082-32-98","Password":"IbCtjOg","Address":"Oak Terrace 91"}
	{"Id":158,"Name":"Willie Olson","Username":"8Stephens","Email":"bMason@Aimbo.com","Phone":"5-028-673-99-90","Password":"nXDQ2w0h","Address":"Jenifer Crossing 95"}
	{"Id":159,"Name":"Jose Stanley Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"vBailey","Email":"FrankPierce@Jetwire.mil","Phone":"7-196-378-63-96","Password":"KnLT67","Address":"Waxwing Park 82"}
	{"Id":160,"Name":"Frank Harris","Username":"BobbyOlson","Email":"JamesBishop@Avamba.mil","Phone":"717-03-41","Password":"5vY2712n2","Address":"Dahle Alley 38"}
	{"Id":161,"Name":"Shawn Perry","Username":"lLawson","Email":"zMedina@Mydeo.com","Phone":"2-151-202-72-15","Password":"X3yOWIG","Address":"Kedzie Park 76"}
	{"Id":162,"Name":"Lois Mendoza","Username":"vel_sunt","Email":"ShawnWheeler@Jaxworks.gov","Phone":"612-67-81","Password":"94tBh6","Address":"Corscot Crossing 55"}
	{"Id":163,"Name":"Mrs. Ms. Miss Mary Knight","Username":"amet_ratione","Email":"rerum_placeat_non@Feedmix.gov","Phone":"053-91-18","Password":"2Zf9eJYZ8f","Address":"Blaine Alley 65"}
	{"Id":164,"Name":"Angela Harrison","Username":"est_non","Email":"zYoung@Realblab.biz","Phone":"9-589-128-64-05","Password":"VIDwt6PAfTU","Address":"Reinke Street 64"}
	{"Id":165,"Name":"Jacqueline Watson","Username":"JeremyReyes","Email":"ChristopherCruz@Twitterlist.edu","Phone":"616-86-83","Password":"HT3ld20r","Address":"Surrey Lane 71"}
	{"Id":166,"Name":"Thomas Tucker","Username":"cDunn","Email":"fBarnes@Topiclounge.gov","Phone":"440-08-40","Password":"s9aeIzb","Address":"Ilene Alley 97"}
	{"Id":167,"Name":"Joe Palmer Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"oGardner","Email":"yHughes@Lazzy.biz","Phone":"9-128-621-64-28","Password":"CDmEy24S","Address":"Messerschmidt Court 13"}
	{"Id":168,"Name":"Philip Boyd","Username":"mWagner","Email":"qui_nisi_voluptate@Voomm.biz","Phone":"962-98-36","Password":"yRAC7R","Address":"Independence Place 57"}
	{"Id":169,"Name":"Marie Larson","Username":"JudyCruz","Email":"wWard@Youopia.name","Phone":"6-404-061-79-36","Password":"uE4BGhXI","Address":"Carpenter Court 41"}
	{"Id":170,"Name":"Walter Bailey","Username":"quia_esse_quos","Email":"MatthewBrown@Jaxbean.gov","Phone":"018-49-10","Password":"JF89gbUYB","Address":"Aberg Lane 43"}
	{"Id":171,"Name":"Johnny Flores Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"2Ramirez","Email":"hRichards@Thoughtstorm.gov","Phone":"7-997-421-68-96","Password":"9nnMLqcZk","Address":"Scofield Street 72"}
	{"Id":172,"Name":"Mrs. Ms. Miss Emily Fox","Username":"EmilyBennett","Email":"JeanMorrison@Realmix.biz","Phone":"0-015-728-62-75","Password":"WNVFFfUgWXex","Address":"Waubesa Crossing 92"}
	{"Id":173,"Name":"Jeffrey Price","Username":"rStevens","Email":"AmandaMurphy@Riffpath.edu","Phone":"1-156-787-12-59","Password":"h8geYFK96","Address":"Hagan Hill 88"}
	{"Id":174,"Name":"Lori Roberts","Username":"praesentium","Email":"wHughes@Omba.org","Phone":"246-73-20","Password":"0AIY41jbjCUv","Address":"Cascade Drive 90"}
	{"Id":175,"Name":"Julie Arnold I II III IV V MD DDS PhD DVM","Username":"HowardDunn","Email":"9Reed@Shufflebeat.mil","Phone":"346-78-95","Password":"MaWU4ag1","Address":"Washington Parkway 36"}
	{"Id":176,"Name":"Lori Scott I II III IV V MD DDS PhD DVM","Username":"cRuiz","Email":"PaulHolmes@Oyonder.info","Phone":"1-639-016-15-62","Password":"pmuSkv24","Address":"Iowa Circle 65"}
	{"Id":177,"Name":"Kimberly Hansen","Username":"rStone","Email":"et_rerum@Youbridge.gov","Phone":"3-066-228-43-41","Password":"LSL49HSZDFTQ","Address":"Charing Cross Way 17"}
	{"Id":178,"Name":"Jack Richardson","Username":"mollitia","Email":"voluptatum_molestias@Cogilith.org","Phone":"143-45-89","Password":"9iEbRnvKk","Address":"Michigan Circle 22"}
	{"Id":179,"Name":"Jeffrey James","Username":"hMurphy","Email":"kPerry@Yata.biz","Phone":"4-267-716-81-76","Password":"c1nJuBld0H4","Address":"Porter Avenue 56"}
	{"Id":180,"Name":"Mrs. Ms. Miss Andrea Alvarez","Username":"jLittle","Email":"1Howell@Dabjam.name","Phone":"059-43-05","Password":"puCcreMHkI","Address":"Jay Street 79"}
	{"Id":181,"Name":"Ashley Olson I II III IV V MD DDS PhD DVM","Username":"pMatthews","Email":"aut_ut@Realcube.com","Phone":"953-25-35","Password":"3BcCBO","Address":"Grasskamp Circle 76"}
	{"Id":182,"Name":"Pamela Nelson","Username":"LisaWells","Email":"eBowman@Realbridge.biz","Phone":"7-591-939-41-40","Password":"VquvNQ4rvI","Address":"Reindahl Junction 10"}
	{"Id":183,"Name":"Antonio Rivera","Username":"SteveHamilton","Email":"eligendi_perspiciatis@Zoonder.net","Phone":"1-745-503-19-83","Password":"X8xI3KMcmY","Address":"Elka Drive 89"}
	{"Id":184,"Name":"Patricia Frazier","Username":"commodi","Email":"RobertDixon@Divanoodle.mil","Phone":"956-18-19","Password":"Naitw3Z","Address":"Mayer Terrace 79"}
	{"Id":185,"Name":"Mr. Dr. Philip Simpson","Username":"8Tucker","Email":"dignissimos@Zoonder.org","Phone":"0-122-758-44-00","Password":"05K9cU","Address":"Annamark Court 67"}
	{"Id":186,"Name":"Russell Campbell","Username":"DavidMatthews","Email":"voluptatum_odit_rerum@Dabshots.info","Phone":"0-501-398-20-95","Password":"KMbwO6DtMl","Address":"Northview Drive 12"}
	{"Id":187,"Name":"Chris Hicks","Username":"3Fields","Email":"wRobinson@Shuffledrive.gov","Phone":"139-97-32","Password":"zKTjbv1Nk","Address":"Redwing Junction 39"}
	{"Id":188,"Name":"Christine Williamson","Username":"BillyLee","Email":"LawrenceScott@Voonyx.biz","Phone":"6-412-508-26-10","Password":"lsYP4y6PvA","Address":"Melvin Road 56"}
	{"Id":189,"Name":"Melissa Alvarez","Username":"animi_harum","Email":"sit@Feedspan.net","Phone":"721-42-09","Password":"gTETQDkwP1o","Address":"Iowa Lane 78"}
	{"Id":190,"Name":"Jason Fernandez","Username":"accusantium_laudantium_et","Email":"RussellWood@Jamia.com","Phone":"3-952-117-47-78","Password":"EUmHSrkvqmB6","Address":"Miller Point 79"}
	{"Id":191,"Name":"Mr. Dr. Walter Dean","Username":"et","Email":"HarryKing@Oyoloo.info","Phone":"9-633-583-99-70","Password":"9a8erb","Address":"Ilene Park 82"}
	{"Id":192,"Name":"Philip Ruiz","Username":"aHansen","Email":"ChristinaJordan@Riffpath.net","Phone":"752-10-60","Password":"w2S0fi","Address":"Bowman Court 92"}
	{"Id":193,"Name":"Brian Sims","Username":"fDixon","Email":"lDixon@Zoomcast.name","Phone":"6-653-049-59-29","Password":"3wd2Fr","Address":"Blue Bill Park Junction 77"}
	{"Id":194,"Name":"Steve Castillo","Username":"vArmstrong","Email":"ab_facilis@Blogpad.org","Phone":"9-967-484-33-43","Password":"zM7zRPU","Address":"Russell Crossing 11"}
	{"Id":195,"Name":"Laura Porter I II III IV V MD DDS PhD DVM","Username":"consequatur_cumque","Email":"1Barnes@Zoombeat.name","Phone":"3-414-774-33-02","Password":"IpAFVJwtM","Address":"Sutteridge Street 3"}
	{"Id":196,"Name":"Juan Medina","Username":"GaryHunt","Email":"facilis_asperiores@Jabberbean.edu","Phone":"4-692-174-70-77","Password":"1X3iv92mR","Address":"Jay Center 98"}
	{"Id":197,"Name":"Jack Rice","Username":"CraigBurns","Email":"JuanLopez@Kayveo.biz","Phone":"7-481-140-09-92","Password":"XNpJnSWn","Address":"Killdeer Lane 68"}
	{"Id":198,"Name":"Diane Duncan","Username":"AliceRay","Email":"7Howell@Rhyloo.com","Phone":"0-526-790-68-97","Password":"2OtQCFw","Address":"Shoshone Parkway 27"}
	{"Id":199,"Name":"Mrs. Ms. Miss Virginia Snyder","Username":"sFord","Email":"zMcdonald@Aimbo.edu","Phone":"7-881-230-06-92","Password":"AumLxsJRByeV","Address":"High Crossing Park 85"}
	{"Id":200,"Name":"Juan Peters","Username":"soluta","Email":"aut_explicabo_voluptatum@Cogibox.com","Phone":"4-732-469-97-06","Password":"hqTPJCx4kCdD","Address":"Moose Avenue 3"}
	{"Id":201,"Name":"Russell Sullivan","Username":"quisquam_maxime_dicta","Email":"BrendaRamos@Skyndu.mil","Phone":"0-414-276-17-25","Password":"Ky2JlFR","Address":"Ridgeway Avenue 67"}
	{"Id":202,"Name":"Gregory Garcia","Username":"soluta_asperiores_et","Email":"RobertGarza@Jabbersphere.net","Phone":"9-213-262-52-03","Password":"p1bCrpeMaf","Address":"Steensland Road 14"}
	{"Id":203,"Name":"Mr. Dr. Phillip Gray","Username":"GloriaBrown","Email":"JonathanJacobs@Wikivu.biz","Phone":"830-87-56","Password":"17nQExYhKWmO","Address":"Garrison Pass 47"}
	{"Id":204,"Name":"Alice Cunningham","Username":"8Greene","Email":"non_quibusdam_non@Topiczoom.org","Phone":"757-85-91","Password":"laADyDD","Address":"Ramsey Junction 43"}
	{"Id":205,"Name":"Anna Butler","Username":"quasi_ipsa","Email":"aut_ut_omnis@Zoombox.name","Phone":"4-338-004-58-52","Password":"s64SPZg0d5","Address":"Ronald Regan Street 73"}
	{"Id":206,"Name":"Tammy Perkins","Username":"eveniet_ut_iusto","Email":"3Anderson@Blogtag.edu","Phone":"7-770-520-07-66","Password":"WOQlALmk","Address":"Fairview Plaza 28"}
	{"Id":207,"Name":"Karen Simmons","Username":"KennethWheeler","Email":"fMason@Zoombeat.biz","Phone":"7-822-095-69-86","Password":"7lDxujE5xsfZ","Address":"Spenser Trail 62"}
	{"Id":208,"Name":"Julia Weaver","Username":"cWallace","Email":"MildredBanks@Leenti.mil","Phone":"719-36-61","Password":"DgzAL6FC0","Address":"Sunnyside Junction 99"}
	{"Id":209,"Name":"Brenda Griffin I II III IV V MD DDS PhD DVM","Username":"EmilyCollins","Email":"zCook@Voolith.biz","Phone":"1-022-956-18-46","Password":"ig7rNvd","Address":"Rutledge Point 60"}
	{"Id":210,"Name":"Mrs. Ms. Miss Jessica Ferguson","Username":"impedit","Email":"2Wheeler@Meezzy.gov","Phone":"3-430-883-10-99","Password":"oWlnXDwdaNg","Address":"Graceland Junction 10"}
	{"Id":211,"Name":"Louise Riley","Username":"eBanks","Email":"StephenNichols@Pixonyx.mil","Phone":"3-342-361-87-23","Password":"K9PisaE0tVL","Address":"Homewood Crossing 5"}
	{"Id":212,"Name":"Timothy Murphy","Username":"2Payne","Email":"DeniseYoung@Agimba.org","Phone":"0-529-668-36-77","Password":"4AEZ4119Y1J","Address":"Sycamore Drive 7"}
	{"Id":213,"Name":"Mrs. Ms. Miss Debra King","Username":"2Montgomery","Email":"qui_accusantium_sint@Babbleblab.net","Phone":"227-71-49","Password":"2dJlNd","Address":"Lillian Way 87"}
	{"Id":214,"Name":"Roy Harper","Username":"GloriaRichards","Email":"dolorem_animi@Eire.info","Phone":"6-710-485-46-50","Password":"08NJZeWQ","Address":"Scoville Court 8"}
	{"Id":215,"Name":"Todd Fernandez","Username":"aut","Email":"BeverlyPerkins@Feedmix.info","Phone":"6-707-147-90-82","Password":"IMUajtfYu","Address":"Canary Drive 44"}
	{"Id":216,"Name":"Annie Allen","Username":"zMcdonald","Email":"mDaniels@Twitterwire.mil","Phone":"868-88-24","Password":"gXlpGW6f6","Address":"Alpine Terrace 87"}
	{"Id":217,"Name":"Maria Nguyen","Username":"GregoryRice","Email":"NicholasOrtiz@Eazzy.name","Phone":"7-074-079-60-84","Password":"0wf3Bpd74","Address":"Green Ridge Center 24"}
	{"Id":218,"Name":"Kimberly Garza","Username":"fGutierrez","Email":"hic_ex@Aibox.edu","Phone":"3-017-878-04-15","Password":"DILxnbH0Qtwj","Address":"Springs Point 87"}
	{"Id":219,"Name":"Craig Spencer","Username":"gBrooks","Email":"dolores_incidunt_illum@Zoomdog.com","Phone":"242-01-60","Password":"tVD9NNEFGeP","Address":"Barnett Street 60"}
	{"Id":220,"Name":"Gerald Woods Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"wFlores","Email":"lKelley@Kayveo.mil","Phone":"0-045-028-38-53","Password":"QvUenbt","Address":"Dorton Street 59"}
	{"Id":221,"Name":"Craig Bailey","Username":"facere_ratione","Email":"LoriStanley@Avamba.info","Phone":"851-53-44","Password":"uZFRxRs","Address":"Service Junction 46"}
	{"Id":222,"Name":"Evelyn Owens","Username":"SharonWeaver","Email":"aut_officiis@Thoughtblab.mil","Phone":"6-387-615-63-03","Password":"vBwnCZI7Xf","Address":"5th Crossing 75"}
	{"Id":223,"Name":"Victor Howell","Username":"LillianSanders","Email":"vMorales@Quaxo.biz","Phone":"2-662-911-34-00","Password":"mwooQBNN","Address":"American Ash Terrace 8"}
	{"Id":224,"Name":"Mr. Dr. Todd Hicks","Username":"culpa_id","Email":"et_reprehenderit@Jazzy.org","Phone":"562-66-01","Password":"Ig3aXE6yySi","Address":"Oak Center 99"}
	{"Id":225,"Name":"Catherine Carroll","Username":"mWheeler","Email":"harum_sequi_et@Feedmix.org","Phone":"387-94-88","Password":"E5Dw7gtuQCM","Address":"Hagan Avenue 8"}
	{"Id":226,"Name":"Deborah Perkins","Username":"MatthewWhite","Email":"AaronHamilton@Thoughtstorm.com","Phone":"2-046-137-66-04","Password":"N7JkWPHM0R","Address":"Columbus Park 30"}
	{"Id":227,"Name":"Timothy Kim","Username":"modi","Email":"zRuiz@Viva.mil","Phone":"7-345-893-58-95","Password":"bQC0nfyp","Address":"Maple Wood Parkway 65"}
	{"Id":228,"Name":"Janice Harris","Username":"AshleyWard","Email":"4Gordon@Devpulse.info","Phone":"6-250-045-39-99","Password":"io48sVU","Address":"Blaine Crossing 28"}
	{"Id":229,"Name":"Robin Crawford","Username":"2Reynolds","Email":"BruceGutierrez@Mudo.org","Phone":"4-350-093-32-61","Password":"BaDQ0O","Address":"Waywood Center 53"}
	{"Id":230,"Name":"Larry Anderson","Username":"1Reed","Email":"3Stanley@Fivespan.biz","Phone":"727-68-04","Password":"QMzHAw93JNn","Address":"Division Road 44"}
	{"Id":231,"Name":"Benjamin Lane","Username":"omnis_perferendis_corrupti","Email":"GeorgeLawson@Meembee.gov","Phone":"716-91-54","Password":"zCS7HHtC","Address":"Hoepker Point 74"}
	{"Id":232,"Name":"Karen Bailey","Username":"mHarris","Email":"et@Kwimbee.com","Phone":"259-94-67","Password":"TrJXyf","Address":"Twin Pines Alley 93"}
	{"Id":233,"Name":"Sean Nelson","Username":"3Morales","Email":"qui_voluptatum@Topiclounge.biz","Phone":"031-49-58","Password":"gaBCcKt","Address":"Butternut Alley 30"}
	{"Id":234,"Name":"Albert Reid","Username":"JaniceMartin","Email":"2Oliver@Thoughtbeat.name","Phone":"795-49-36","Password":"cLe9RtdSPiiF","Address":"Old Gate Avenue 16"}
	{"Id":235,"Name":"Lori Meyer","Username":"JustinMorales","Email":"uGibson@Gevee.mil","Phone":"088-32-01","Password":"BXoEnWWhus","Address":"Holy Cross Trail 60"}
	{"Id":236,"Name":"Robin Holmes","Username":"9Perry","Email":"KarenShaw@Voolith.info","Phone":"687-42-35","Password":"lOvKpm1","Address":"Village Road 50"}
	{"Id":237,"Name":"Mary Roberts","Username":"eos_eaque_totam","Email":"reiciendis@Jaloo.org","Phone":"543-16-62","Password":"HgigFQrp","Address":"Fisk Crossing 91"}
	{"Id":238,"Name":"Scott Simpson","Username":"expedita","Email":"uLarson@LiveZ.info","Phone":"5-481-171-05-30","Password":"trg6b4","Address":"Cordelia Avenue 46"}
	{"Id":239,"Name":"Jeremy Perez","Username":"quisquam_eos_voluptate","Email":"consequuntur@Vipe.mil","Phone":"542-82-15","Password":"68plO8Jys","Address":"Tony Center 99"}
	{"Id":240,"Name":"Kathy Fuller","Username":"rerum_sunt_quisquam","Email":"dolores_fuga@Lazz.gov","Phone":"4-334-000-09-06","Password":"j5UeSD9eagV","Address":"Oakridge Street 76"}
	{"Id":241,"Name":"Donna Mitchell I II III IV V MD DDS PhD DVM","Username":"JoshuaSanders","Email":"quam_aspernatur@Meevee.com","Phone":"303-14-90","Password":"cRIHvVQfT","Address":"Bultman Point 60"}
	{"Id":242,"Name":"Phillip Torres","Username":"yCruz","Email":"illum@Twinder.name","Phone":"1-826-912-31-59","Password":"eqtsQns2abXw","Address":"Morning Center 81"}
	{"Id":243,"Name":"Jesse Gordon","Username":"JoseHoward","Email":"assumenda_quia@Topiczoom.com","Phone":"9-961-610-93-02","Password":"3gvf7SNeF","Address":"Killdeer Drive 21"}
	{"Id":244,"Name":"Ralph Carroll","Username":"EvelynJackson","Email":"doloremque@Devify.com","Phone":"639-01-40","Password":"BnfKiu","Address":"Clemons Center 86"}
	{"Id":245,"Name":"Jack Wallace","Username":"RalphFisher","Email":"yCarr@Ntag.name","Phone":"2-421-346-82-45","Password":"hJn8ol","Address":"Sunnyside Court 91"}
	{"Id":246,"Name":"Sean Allen","Username":"quas_eius_dolor","Email":"uHarper@Plajo.com","Phone":"0-168-355-24-37","Password":"skOzB0vWM9GI","Address":"Alpine Crossing 25"}
	{"Id":247,"Name":"Kathleen Hawkins","Username":"6Davis","Email":"WalterMartin@Skajo.com","Phone":"6-969-667-68-85","Password":"jFgSg9rs1G","Address":"Hagan Trail 83"}
	{"Id":248,"Name":"Susan Fernandez","Username":"zHicks","Email":"qui_assumenda_voluptatem@Skynoodle.com","Phone":"7-449-180-44-53","Password":"FxUqxwsMQoJH","Address":"Dryden Center 90"}
	{"Id":249,"Name":"Michael Robertson","Username":"AndrewRuiz","Email":"gEvans@Realcube.mil","Phone":"5-481-294-58-53","Password":"YMir1dYyvBR","Address":"Butternut Park 29"}
	{"Id":250,"Name":"Jesse Willis","Username":"HowardPerez","Email":"2Montgomery@Jaxnation.mil","Phone":"1-235-554-63-04","Password":"vDfqUDEdcF","Address":"Petterle Point 27"}
	{"Id":251,"Name":"Harold Watkins","Username":"StevenGrant","Email":"doloremque@Skyba.org","Phone":"785-49-56","Password":"6Xs4IZd5c4lJ","Address":"Artisan Trail 99"}
	{"Id":252,"Name":"Joan Reid","Username":"CarlBaker","Email":"rChavez@Realcube.gov","Phone":"0-376-259-99-75","Password":"mFZ1boD","Address":"Rusk Alley 74"}
	{"Id":253,"Name":"Mr. Dr. Larry Dean","Username":"bRyan","Email":"enim_ad_et@Jetwire.com","Phone":"159-47-10","Password":"OjDKxs4cN","Address":"Evergreen Trail 95"}
	{"Id":254,"Name":"Mr. Dr. Peter Howell","Username":"dolor","Email":"KellyWelch@Voolith.biz","Phone":"698-48-56","Password":"adIKseC","Address":"Mifflin Way 68"}
	{"Id":255,"Name":"Brenda Kennedy","Username":"ThomasCarr","Email":"unde_dignissimos_voluptas@Wordtune.com","Phone":"6-763-434-83-28","Password":"rtAE7vq","Address":"Moose Street 33"}
	{"Id":256,"Name":"Mr. Dr. Billy Reed","Username":"optio_suscipit","Email":"dMatthews@Flipstorm.edu","Phone":"2-122-548-76-32","Password":"VpfTlzq6","Address":"Vermont Parkway 19"}
	{"Id":257,"Name":"Roger Bowman","Username":"sAllen","Email":"et_doloremque@Skyba.biz","Phone":"6-396-218-91-39","Password":"fz3De5","Address":"Corscot Point 47"}
	{"Id":258,"Name":"Lisa Diaz","Username":"vero_voluptatibus_fugiat","Email":"8Allen@Layo.mil","Phone":"192-50-79","Password":"bJNJGSXA","Address":"Comanche Park 63"}
	{"Id":259,"Name":"Carol Sims","Username":"3Thompson","Email":"in@Katz.org","Phone":"761-73-82","Password":"WzMHkSrv","Address":"Bellgrove Plaza 13"}
	{"Id":260,"Name":"Thomas Fuller","Username":"AlicePrice","Email":"et_dolorem@Podcat.gov","Phone":"2-516-310-45-09","Password":"8CQTk1NwXCh1","Address":"Victoria Crossing 24"}
	{"Id":261,"Name":"Brandon Thomas","Username":"est_neque_dignissimos","Email":"et_et_aut@Yodoo.info","Phone":"367-92-90","Password":"8XmRGbwuvD","Address":"Logan Road 99"}
	{"Id":262,"Name":"Cheryl Hill I II III IV V MD DDS PhD DVM","Username":"ratione","Email":"qui_ad@Wordpedia.edu","Phone":"9-624-889-62-76","Password":"g5mWqRfzBj","Address":"Barby Court 19"}
	{"Id":263,"Name":"Joan Freeman","Username":"JosephBoyd","Email":"nihil_iusto@Oozz.biz","Phone":"2-263-762-08-84","Password":"ovTJVH","Address":"Westridge Court 80"}
	{"Id":264,"Name":"Stephanie Fuller","Username":"tempora_minima","Email":"cHenry@Innotype.biz","Phone":"138-28-26","Password":"p7gjr3","Address":"Banding Lane 23"}
	{"Id":265,"Name":"George Spencer","Username":"consequatur","Email":"tFrazier@Quinu.name","Phone":"1-702-215-87-01","Password":"v3KfSmu","Address":"Meadow Ridge Parkway 17"}
	{"Id":266,"Name":"Heather Arnold","Username":"JimmyHayes","Email":"qui_voluptas_vero@Lajo.mil","Phone":"6-204-820-02-34","Password":"ZoE0pNkuvnfO","Address":"Burrows Street 65"}
	{"Id":267,"Name":"Christopher Nguyen","Username":"aut","Email":"RachelMurray@Skipfire.edu","Phone":"603-97-80","Password":"VxEsSPl","Address":"Mifflin Alley 64"}
	{"Id":268,"Name":"Stephanie Ross","Username":"KeithSchmidt","Email":"quam@Thoughtsphere.info","Phone":"286-85-73","Password":"GszOLelIoe","Address":"Loftsgordon Plaza 97"}
	{"Id":269,"Name":"Mrs. Ms. Miss Shirley Martinez","Username":"rSullivan","Email":"tJohnston@Riffpedia.biz","Phone":"5-194-486-72-57","Password":"Mz9ufc8V","Address":"Westport Parkway 96"}
	{"Id":270,"Name":"Henry Stevens","Username":"voluptas_minus","Email":"gReynolds@Blogtags.name","Phone":"685-28-18","Password":"qzMMWMpW1bs","Address":"Transport Point 51"}
	{"Id":271,"Name":"Linda Watkins","Username":"SeanLopez","Email":"nihil@Quimm.gov","Phone":"5-285-750-05-67","Password":"zMvK4QHX","Address":"Florence Junction 32"}
	{"Id":272,"Name":"Justin Howell","Username":"ScottHart","Email":"iMorris@Babblestorm.edu","Phone":"677-76-75","Password":"dve6LfMp","Address":"Farmco Parkway 58"}
	{"Id":273,"Name":"Rose Brown","Username":"kOliver","Email":"RussellAndrews@Lazzy.org","Phone":"133-64-73","Password":"iEmOex3iNMW","Address":"Manitowish Road 44"}
	{"Id":274,"Name":"Joseph Warren","Username":"tWelch","Email":"StephenCarter@Edgeify.biz","Phone":"2-131-312-23-62","Password":"j7tOTl","Address":"Grasskamp Plaza 61"}
	{"Id":275,"Name":"Ronald Banks","Username":"HenryGeorge","Email":"7Perez@Topdrive.edu","Phone":"908-77-75","Password":"tpIHRTQCo","Address":"Norway Maple Circle 82"}
	{"Id":276,"Name":"Janice Wood","Username":"hCarr","Email":"BettyHansen@Riffwire.info","Phone":"9-378-106-72-01","Password":"UjUdpuF4fahZ","Address":"Hayes Point 74"}
	{"Id":277,"Name":"Howard Chapman","Username":"et","Email":"cumque_sint_voluptatibus@BlogXS.net","Phone":"1-452-092-84-39","Password":"qttw5wFl","Address":"Morrow Center 1"}
	{"Id":278,"Name":"William Bell Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"fuga_voluptatibus","Email":"BruceWard@Kazu.org","Phone":"6-394-334-53-11","Password":"SgFCKZ","Address":"3rd Street 55"}
	{"Id":279,"Name":"Bruce Ramos","Username":"laudantium_omnis","Email":"sHenry@Twitterworks.mil","Phone":"261-80-65","Password":"6pCXYBL","Address":"Granby Circle 21"}
	{"Id":280,"Name":"Anthony Holmes","Username":"RalphDavis","Email":"rerum_tenetur@Livetube.com","Phone":"558-36-08","Password":"r1jTag","Address":"Hudson Center 46"}
	{"Id":281,"Name":"Carlos Torres","Username":"dicta_harum","Email":"JoseMoore@Tavu.net","Phone":"1-227-125-18-08","Password":"JUiGC1","Address":"Shoshone Lane 10"}
	{"Id":282,"Name":"Dorothy Lopez","Username":"9Gutierrez","Email":"nCarr@Blogtags.name","Phone":"927-95-57","Password":"pP5DsZdT","Address":"Golf Course Terrace 33"}
	{"Id":283,"Name":"Phyllis Fernandez","Username":"voluptatem_iure_corporis","Email":"rBaker@Thoughtstorm.net","Phone":"089-62-25","Password":"4d9rxvPekF","Address":"2nd Center 60"}
	{"Id":284,"Name":"William Payne Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"MelissaFerguson","Email":"WayneLittle@Thoughtblab.info","Phone":"610-96-99","Password":"ky8wGgukpzI","Address":"Pearson Pass 40"}
	{"Id":285,"Name":"Denise Nichols","Username":"sit_corporis","Email":"necessitatibus@Twitterbeat.name","Phone":"1-759-113-71-52","Password":"2okPQ9lNs9q","Address":"Parkside Place 91"}
	{"Id":286,"Name":"Ann Reyes","Username":"bGrant","Email":"JoyceMartin@Thoughtsphere.com","Phone":"2-367-440-26-09","Password":"2ysyg7lYsc2","Address":"Muir Place 54"}
	{"Id":287,"Name":"Mrs. Ms. Miss Robin Wells","Username":"JenniferMurray","Email":"0Burton@Browseblab.edu","Phone":"5-899-607-32-20","Password":"jLFrgp","Address":"Fieldstone Trail 91"}
	{"Id":288,"Name":"Kimberly Bennett","Username":"eveniet","Email":"wArmstrong@Thoughtstorm.com","Phone":"2-516-348-54-68","Password":"CV7udJe","Address":"Nova Court 19"}
	{"Id":289,"Name":"Nicole Hanson","Username":"EarlEdwards","Email":"qui_ratione@Einti.mil","Phone":"620-77-80","Password":"hVFgfZKc5FBh","Address":"Schiller Place 43"}
	{"Id":290,"Name":"Harry Young","Username":"bOliver","Email":"LillianCole@Lajo.org","Phone":"5-458-549-46-92","Password":"u0L4xiliY","Address":"International Hill 8"}
	{"Id":291,"Name":"Steven Harrison Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"DeniseHenry","Email":"MarilynRomero@Digitube.name","Phone":"8-111-468-67-20","Password":"eO9342V","Address":"High Crossing Plaza 74"}
	{"Id":292,"Name":"Jerry Jacobs","Username":"BeverlySanders","Email":"qChavez@Oba.org","Phone":"4-656-229-00-49","Password":"bc8l3aCgQs","Address":"Johnson Center 66"}
	{"Id":293,"Name":"Virginia Gonzales","Username":"eKim","Email":"qui_perspiciatis@Wikizz.edu","Phone":"9-219-298-01-99","Password":"CHTIQAuh","Address":"Merchant Center 6"}
	{"Id":294,"Name":"Michael Stephens","Username":"DianeRomero","Email":"xRuiz@Lajo.info","Phone":"016-69-31","Password":"tgV0S1x","Address":"Sunbrook Junction 39"}
	{"Id":295,"Name":"Daniel Rivera","Username":"non","Email":"RubyGray@Skiba.mil","Phone":"8-395-380-13-08","Password":"0K65d9e6qJy6","Address":"Clemons Hill 13"}
	{"Id":296,"Name":"Clarence Rivera","Username":"lFrazier","Email":"aperiam@Quimba.info","Phone":"795-32-91","Password":"nkRlYTJA","Address":"Porter Hill 3"}
	{"Id":297,"Name":"Victor Edwards","Username":"placeat_omnis","Email":"fPierce@JumpXS.name","Phone":"464-12-68","Password":"aPL883N2pl0V","Address":"Artisan Pass 31"}
	{"Id":298,"Name":"Steve Rivera","Username":"eDavis","Email":"DianaPatterson@Voomm.edu","Phone":"404-57-16","Password":"t44Zxbcwqh","Address":"Merchant Crossing 41"}
	{"Id":299,"Name":"Daniel Burke","Username":"CynthiaFranklin","Email":"modi@Quinu.net","Phone":"5-562-435-03-35","Password":"WJjHPKy78","Address":"Dixon Circle 48"}\
	{"Id":300,"Name":"Christopher Castillo","Username":"cMitchell","Email":"adipisci_totam_qui@Brightbean.mil","Phone":"949-81-91","Password":"SRd9QrNg","Address":"Russell Street 66"}
	{"Id":301,"Name":"Mrs. Ms. Miss Jessica Austin","Username":"molestias_eos_deserunt","Email":"itaque@Gabtune.name","Phone":"741-66-99","Password":"IEgM3mUSU64N","Address":"Di Loreto Pass 33"}`

	b.ResetTimer()
	GetDomainStat(bytes.NewBufferString(data), "com")
}

func benchmarkGetDomainStat(b *testing.B) {
	b.StopTimer()
	r, _ := zip.OpenReader("testdata/users.dat.zip")
	defer r.Close()

	data, _ := r.File[0].Open()

	// b.ResetTimer()
	b.StartTimer()
	GetDomainStat(data, "biz")
	b.StopTimer()
}

// go test -v -count=1 -timeout=30s -tags bench .
func TestGetDomainStat_Time_And_Memory(t *testing.T) {
	bench := func(b *testing.B) {
		b.StopTimer()

		r, err := zip.OpenReader("testdata/users.dat.zip")
		require.NoError(t, err)
		defer r.Close()

		require.Equal(t, 1, len(r.File))

		data, err := r.File[0].Open()
		require.NoError(t, err)

		b.StartTimer()
		stat, err := GetDomainStat(data, "biz")
		b.StopTimer()
		require.NoError(t, err)

		require.Equal(t, expectedBizStat, stat)
	}

	result := testing.Benchmark(bench)
	mem := result.MemBytes
	t.Logf("time used: %s", result.T)
	t.Logf("memory used: %dMb", mem/mb)

	require.Less(t, int64(result.T), int64(timeLimit), "the program is too slow")
	require.Less(t, mem, memoryLimit, "the program is too greedy")
}

var expectedBizStat = DomainStat{
	"abata.biz":         25,
	"abatz.biz":         25,
	"agimba.biz":        28,
	"agivu.biz":         17,
	"aibox.biz":         31,
	"ailane.biz":        23,
	"aimbo.biz":         25,
	"aimbu.biz":         36,
	"ainyx.biz":         35,
	"aivee.biz":         25,
	"avamba.biz":        21,
	"avamm.biz":         17,
	"avavee.biz":        35,
	"avaveo.biz":        30,
	"babbleblab.biz":    29,
	"babbleopia.biz":    36,
	"babbleset.biz":     28,
	"babblestorm.biz":   29,
	"blognation.biz":    32,
	"blogpad.biz":       34,
	"blogspan.biz":      21,
	"blogtag.biz":       23,
	"blogtags.biz":      34,
	"blogxs.biz":        35,
	"bluejam.biz":       36,
	"bluezoom.biz":      27,
	"brainbox.biz":      30,
	"brainlounge.biz":   38,
	"brainsphere.biz":   31,
	"brainverse.biz":    39,
	"brightbean.biz":    23,
	"brightdog.biz":     32,
	"browseblab.biz":    31,
	"browsebug.biz":     25,
	"browsecat.biz":     34,
	"browsedrive.biz":   24,
	"browsetype.biz":    34,
	"browsezoom.biz":    29,
	"bubblebox.biz":     19,
	"bubblemix.biz":     38,
	"bubbletube.biz":    34,
	"buzzbean.biz":      26,
	"buzzdog.biz":       30,
	"buzzshare.biz":     26,
	"buzzster.biz":      28,
	"camido.biz":        27,
	"camimbo.biz":       36,
	"centidel.biz":      32,
	"centimia.biz":      17,
	"centizu.biz":       18,
	"chatterbridge.biz": 30,
	"chatterpoint.biz":  32,
	"cogibox.biz":       30,
	"cogidoo.biz":       34,
	"cogilith.biz":      24,
	"dabfeed.biz":       26,
	"dabjam.biz":        30,
	"dablist.biz":       30,
	"dabshots.biz":      33,
	"dabtype.biz":       21,
	"dabvine.biz":       26,
	"dabz.biz":          19,
	"dazzlesphere.biz":  24,
	"demimbu.biz":       27,
	"demivee.biz":       39,
	"demizz.biz":        30,
	"devbug.biz":        20,
	"devcast.biz":       35,
	"devify.biz":        27,
	"devpoint.biz":      26,
	"devpulse.biz":      27,
	"devshare.biz":      30,
	"digitube.biz":      30,
	"divanoodle.biz":    33,
	"divape.biz":        32,
	"divavu.biz":        28,
	"dynabox.biz":       66,
	"dynava.biz":        21,
	"dynazzy.biz":       29,
	"eabox.biz":         28,
	"eadel.biz":         25,
	"eamia.biz":         18,
	"eare.biz":          30,
	"eayo.biz":          30,
	"eazzy.biz":         27,
	"edgeblab.biz":      29,
	"edgeclub.biz":      29,
	"edgeify.biz":       36,
	"edgepulse.biz":     21,
	"edgetag.biz":       24,
	"edgewire.biz":      29,
	"eidel.biz":         33,
	"eimbee.biz":        22,
	"einti.biz":         19,
	"eire.biz":          28,
	"fadeo.biz":         35,
	"fanoodle.biz":      23,
	"fatz.biz":          30,
	"feedbug.biz":       29,
	"feedfire.biz":      30,
	"feedfish.biz":      35,
	"feedmix.biz":       31,
	"feednation.biz":    24,
	"feedspan.biz":      28,
	"fivebridge.biz":    20,
	"fivechat.biz":      29,
	"fiveclub.biz":      23,
	"fivespan.biz":      27,
	"flashdog.biz":      20,
	"flashpoint.biz":    35,
	"flashset.biz":      30,
	"flashspan.biz":     32,
	"flipbug.biz":       27,
	"flipopia.biz":      30,
	"flipstorm.biz":     21,
	"fliptune.biz":      29,
	"gabcube.biz":       29,
	"gabspot.biz":       24,
	"gabtune.biz":       29,
	"gabtype.biz":       29,
	"gabvine.biz":       24,
	"geba.biz":          24,
	"gevee.biz":         23,
	"gigabox.biz":       28,
	"gigaclub.biz":      25,
	"gigashots.biz":     26,
	"gigazoom.biz":      29,
	"innojam.biz":       26,
	"innotype.biz":      27,
	"innoz.biz":         24,
	"izio.biz":          26,
	"jabberbean.biz":    28,
	"jabbercube.biz":    31,
	"jabbersphere.biz":  55,
	"jabberstorm.biz":   22,
	"jabbertype.biz":    27,
	"jaloo.biz":         35,
	"jamia.biz":         33,
	"janyx.biz":         33,
	"jatri.biz":         18,
	"jaxbean.biz":       28,
	"jaxnation.biz":     21,
	"jaxspan.biz":       27,
	"jaxworks.biz":      30,
	"jayo.biz":          44,
	"jazzy.biz":         32,
	"jetpulse.biz":      25,
	"jetwire.biz":       26,
	"jumpxs.biz":        29,
	"kamba.biz":         30,
	"kanoodle.biz":      19,
	"kare.biz":          30,
	"katz.biz":          62,
	"kaymbo.biz":        34,
	"kayveo.biz":        22,
	"kazio.biz":         21,
	"kazu.biz":          16,
	"kimia.biz":         25,
	"kwideo.biz":        17,
	"kwilith.biz":       25,
	"kwimbee.biz":       34,
	"kwinu.biz":         15,
	"lajo.biz":          20,
	"latz.biz":          24,
	"layo.biz":          32,
	"lazz.biz":          27,
	"lazzy.biz":         26,
	"leenti.biz":        26,
	"leexo.biz":         32,
	"linkbridge.biz":    38,
	"linkbuzz.biz":      24,
	"linklinks.biz":     31,
	"linktype.biz":      31,
	"livefish.biz":      31,
	"livepath.biz":      23,
	"livetube.biz":      53,
	"livez.biz":         28,
	"meedoo.biz":        23,
	"meejo.biz":         24,
	"meembee.biz":       26,
	"meemm.biz":         23,
	"meetz.biz":         33,
	"meevee.biz":        62,
	"meeveo.biz":        27,
	"meezzy.biz":        24,
	"miboo.biz":         26,
	"midel.biz":         28,
	"minyx.biz":         25,
	"mita.biz":          29,
	"mudo.biz":          36,
	"muxo.biz":          25,
	"mybuzz.biz":        32,
	"mycat.biz":         32,
	"mydeo.biz":         20,
	"mydo.biz":          30,
	"mymm.biz":          21,
	"mynte.biz":         54,
	"myworks.biz":       27,
	"nlounge.biz":       25,
	"npath.biz":         33,
	"ntag.biz":          28,
	"ntags.biz":         32,
	"oba.biz":           22,
	"oloo.biz":          19,
	"omba.biz":          26,
	"ooba.biz":          27,
	"oodoo.biz":         30,
	"oozz.biz":          22,
	"oyoba.biz":         27,
	"oyoloo.biz":        30,
	"oyonder.biz":       29,
	"oyondu.biz":        23,
	"oyope.biz":         24,
	"oyoyo.biz":         32,
	"ozu.biz":           18,
	"photobean.biz":     25,
	"photobug.biz":      57,
	"photofeed.biz":     25,
	"photojam.biz":      35,
	"photolist.biz":     19,
	"photospace.biz":    33,
	"pixoboo.biz":       14,
	"pixonyx.biz":       30,
	"pixope.biz":        32,
	"plajo.biz":         32,
	"plambee.biz":       29,
	"podcat.biz":        31,
	"quamba.biz":        31,
	"quatz.biz":         54,
	"quaxo.biz":         25,
	"quimba.biz":        25,
	"quimm.biz":         33,
	"quinu.biz":         60,
	"quire.biz":         25,
	"realblab.biz":      32,
	"realbridge.biz":    30,
	"realbuzz.biz":      22,
	"realcube.biz":      57,
	"realfire.biz":      37,
	"reallinks.biz":     25,
	"realmix.biz":       27,
	"realpoint.biz":     22,
	"rhybox.biz":        30,
	"rhycero.biz":       28,
	"rhyloo.biz":        32,
	"rhynoodle.biz":     25,
	"rhynyx.biz":        17,
	"rhyzio.biz":        36,
	"riffpath.biz":      21,
	"riffpedia.biz":     33,
	"riffwire.biz":      31,
	"roodel.biz":        29,
	"roombo.biz":        29,
	"roomm.biz":         32,
	"rooxo.biz":         34,
	"shufflebeat.biz":   32,
	"shuffledrive.biz":  25,
	"shufflester.biz":   26,
	"shuffletag.biz":    23,
	"skaboo.biz":        35,
	"skajo.biz":         26,
	"skalith.biz":       30,
	"skiba.biz":         22,
	"skibox.biz":        27,
	"skidoo.biz":        24,
	"skilith.biz":       29,
	"skimia.biz":        45,
	"skinder.biz":       25,
	"skinix.biz":        23,
	"skinte.biz":        39,
	"skipfire.biz":      29,
	"skippad.biz":       26,
	"skipstorm.biz":     30,
	"skiptube.biz":      26,
	"skivee.biz":        34,
	"skyba.biz":         40,
	"skyble.biz":        32,
	"skyndu.biz":        32,
	"skynoodle.biz":     28,
	"skyvu.biz":         34,
	"snaptags.biz":      33,
	"tagcat.biz":        33,
	"tagchat.biz":       37,
	"tagfeed.biz":       30,
	"tagopia.biz":       17,
	"tagpad.biz":        28,
	"tagtune.biz":       22,
	"talane.biz":        22,
	"tambee.biz":        24,
	"tanoodle.biz":      38,
	"tavu.biz":          37,
	"tazz.biz":          27,
	"tazzy.biz":         28,
	"tekfly.biz":        31,
	"teklist.biz":       26,
	"thoughtbeat.biz":   30,
	"thoughtblab.biz":   24,
	"thoughtbridge.biz": 30,
	"thoughtmix.biz":    33,
	"thoughtsphere.biz": 20,
	"thoughtstorm.biz":  38,
	"thoughtworks.biz":  24,
	"topdrive.biz":      35,
	"topicblab.biz":     32,
	"topiclounge.biz":   21,
	"topicshots.biz":    30,
	"topicstorm.biz":    22,
	"topicware.biz":     35,
	"topiczoom.biz":     38,
	"trilia.biz":        28,
	"trilith.biz":       25,
	"trudeo.biz":        29,
	"trudoo.biz":        28,
	"trunyx.biz":        33,
	"trupe.biz":         34,
	"twimbo.biz":        19,
	"twimm.biz":         30,
	"twinder.biz":       28,
	"twinte.biz":        33,
	"twitterbeat.biz":   33,
	"twitterbridge.biz": 20,
	"twitterlist.biz":   26,
	"twitternation.biz": 22,
	"twitterwire.biz":   21,
	"twitterworks.biz":  39,
	"twiyo.biz":         37,
	"vidoo.biz":         28,
	"vimbo.biz":         21,
	"vinder.biz":        31,
	"vinte.biz":         34,
	"vipe.biz":          25,
	"vitz.biz":          26,
	"viva.biz":          30,
	"voolia.biz":        34,
	"voolith.biz":       26,
	"voomm.biz":         61,
	"voonder.biz":       32,
	"voonix.biz":        32,
	"voonte.biz":        26,
	"voonyx.biz":        25,
	"wikibox.biz":       27,
	"wikido.biz":        21,
	"wikivu.biz":        23,
	"wikizz.biz":        61,
	"wordify.biz":       28,
	"wordpedia.biz":     25,
	"wordtune.biz":      27,
	"wordware.biz":      19,
	"yabox.biz":         24,
	"yacero.biz":        34,
	"yadel.biz":         27,
	"yakidoo.biz":       21,
	"yakijo.biz":        29,
	"yakitri.biz":       26,
	"yambee.biz":        20,
	"yamia.biz":         17,
	"yata.biz":          25,
	"yodel.biz":         26,
	"yodo.biz":          21,
	"yodoo.biz":         24,
	"yombu.biz":         29,
	"yotz.biz":          26,
	"youbridge.biz":     40,
	"youfeed.biz":       32,
	"youopia.biz":       22,
	"youspan.biz":       59,
	"youtags.biz":       22,
	"yoveo.biz":         31,
	"yozio.biz":         33,
	"zava.biz":          29,
	"zazio.biz":         18,
	"zoombeat.biz":      28,
	"zoombox.biz":       30,
	"zoomcast.biz":      38,
	"zoomdog.biz":       29,
	"zoomlounge.biz":    25,
	"zoomzone.biz":      32,
	"zoonder.biz":       29,
	"zoonoodle.biz":     27,
	"zooveo.biz":        22,
	"zoovu.biz":         38,
	"zooxo.biz":         33,
	"zoozzy.biz":        23,
}
