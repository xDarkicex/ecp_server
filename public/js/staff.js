document.addEventListener("DOMContentLoaded", function(e){
    var data = [{
      "Name": "Marvin Rezac",
      "Title": "C.E.O.",
      "Email": "mrezac@easycarewater.com",
      "Phone": "(559) 246-3719",
      "Experience": null,
      "Description": "",
      "City": "Fresno",
      "State": "CA",
      "Country": "United States"
    }, {
      "Name": "Evangelina Serrano",
      "Title": "President",
      "Email": "eserrano@easycarewater.com",
      "Phone": "(559) 246-4480",
      "Experience": null,
      "Description": "",
      "City": "Fresno",
      "State": "CA",
      "Country": "United States"
    }, {
      "Name": "Monica Morgan",
      "Title": "Accountant",
      "Email": "accounting@easycarewater.com",
      "Phone": "(559) 260-9692",
      "Experience": null,
      "Description": "",
      "City": "Fowler",
      "State": "CA",
      "Country": "United States"
    }, {
      "Name": "Gentry Rolofson",
      "Title": "Technomancer",
      "Email": "grolofson@bitdev.io",
      "Phone": "(559) 676-0527",
      "Experience": null,
      "Description": "",
      "City": "Fresno",
      "State": "CA",
      "Country": "United States"
    }, {
      "Name": "Tiffany Rolofson",
      "Title": "Eastern Regional Sales Director, International Sales Director",
      "Email": "trolofson@easycarewater.com",
      "Phone": "(559) 694-1267",
      "Experience": null,
      "Description": "",
      "City": "Fresno",
      "State": "CA",
      "Country": "United States"
    }, {
      "Name": "Rosemarie Arenas",
      "Title": "Western Regional Sales Director, International Sales Director",
      "Email": "rarenas@easycarewater.com",
      "Phone": "(559) 974-8252",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Lucy Baranova",
      "Title": "International Sales Representative",
      "Email": "",
      "Phone": "",
      "Experience": null,
      "Description": "",
      "City": "Athens",
      "State": "",
      "Country": "Greece"
    }, {
      "Name": "Joshua Wyman",
      "Title": "International Sales Representative",
      "Email": "",
      "Phone": "",
      "Experience": null,
      "Description": "",
      "City": "Paris",
      "State": "",
      "Country": "France"
    }, 
    {
      "Name": "Scott Nichols",
      "Title": "Sales Represenative",
      "Email": "snichols@easycarewater.com",
      "Phone": "(469) 401-2130",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Bryan Wiegand",
      "Title": "Sales Representive",
      "Email": "bweigand@easycarewater.com",
      "Phone": "(210) 857-5729",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Todd Wilson",
      "Title": "Senior Sales Represenative",
      "Email": "twilson@easycarewater.com",
      "Phone": "(909) 973-3824",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Matt Wyant",
      "Title": "Sales Represenative",
      "Email": "mwyant@easycarewater.com",
      "Phone": "(623) 738-5061",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Will Bond",
      "Title": "Sales Represenative",
      "Email": "wbond@easycarewater.com",
      "Phone": "(619) 208-4148",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Vicki Brown",
      "Title": "Sales Represenative",
      "Email": "vbrown@easycarewater.com",
      "Phone": "(619) 208-4148",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Bill Hills",
      "Title": "Sales Represenative",
      "Email": "agpoolman@yahoo.com",
      "Phone": "(609) 970-0871",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Kyle Morgan",
      "Title": "Product Specialist",
      "Email": "kmorgan@easycarewater.com",
      "Phone": "(480) 442-4646",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Carlos Mejia",
      "Title": "Sales Represenative",
      "Email": "cmejia@easycarewater.com",
      "Phone": "(786) 369-9903",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Jose Valdovinos",
      "Title": "Sales Represenative",
      "Email": "jvaldovinos@easycarewater.com",
      "Phone": "(818) 581-0991",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Jospeh Mendez",
      "Title": "Plant Manager",
      "Email": "jmendez@easycarewater.com",
      "Phone": "",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Pio Trinidad",
      "Title": "Production Employee",
      "Email": "",
      "Phone": "",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Michael Alvarez",
      "Title": "Shipping Manager",
      "Email": "",
      "Phone": "",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Xavior Mendez",
      "Title": "Packaging",
      "Email": "",
      "Phone": "",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Jimmy Blajos",
      "Title": "Plant Maintenance Manager",
      "Email": "",
      "Phone": "",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Debbie Saldivar",
      "Title": "Operations Manager",
      "Email": "dsaldivar@easycarewater.com",
      "Phone": "(559) 270-7698",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }, {
      "Name": "Jennifer Weaver",
      "Title": "Order Processor",
      "Email": "orderprocessing@easycarewater.com",
      "Phone": "(805) 901-3953",
      "Experience": null,
      "Description": "",
      "City": "",
      "State": "",
      "Country": ""
    }];
    var len = data.length
    for (i = 0; i < len; i++) {
      var container = document.querySelectorAll('.col-lg')[0]
      var info = document.createElement('div')
      if (data[i].Name) {
        var title = document.createElement('h4')
        title.className = "card-title"
        title.innerHTML = data[i].Name
        title.style = "font-weight: bold;"
        info.appendChild(title)
      }
      if (data[i].Title) {
        var jobTitle = document.createElement('p')
        jobTitle.innerHTML = data[i].Title
        info.appendChild(jobTitle)
      }
      if (data[i].Phone) {
        var Phone = document.createElement('p')
        Phone.innerHTML = data[i].Phone
        info.appendChild(Phone)
      }
      if (data[i].Email) {
        var Email = document.createElement('p')
        Email.innerHTML = data[i].Email
        info.appendChild(Email)
      }
       info.className = "card row center"
      container.appendChild(info)
     }
 })

