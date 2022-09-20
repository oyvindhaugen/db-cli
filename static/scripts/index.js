//This accesses the JSON file containing the data from the database
let oXHR = new XMLHttpRequest()
oXHR.onreadystatechange = reportStatus
oXHR.open("GET", "../selectQuery.json", true)
oXHR.send()
function reportStatus() {
    if (oXHR.readyState === 4) {
        createTableFromJSON(this.responseText)
    }
}
//This uses the data from the JSON file to create a dynamic table.
function createTableFromJSON(jsonData) {
    let arrShopping = []
    arrShopping = JSON.parse(jsonData)

    let col = []
    for (let i = 0; i < arrShopping.length; i++) {
        for (let key in arrShopping[i]) {
            if (col.indexOf(key) === -1) {
                col.push(key)
            }
        }
    }

    let table = document.createElement("table")
    let tr = table.insertRow(-1)

    for (let i = 0; i < col.length + 2; i++) {
        let th = document.createElement("th")
        th.innerHTML = col[i]
        tr.appendChild(th)
        if (i === 3) {
            th.innerHTML = 'delBtn'
            tr.appendChild(th)
        }
        if (i === 4) {
            th.innerHTML = 'updtBtn'
            tr.appendChild(th)
        }
    }


    for (let i = 0; i < arrShopping.length; i++) {
        tr = table.insertRow(-1)
        tr.setAttribute('id', "tr" + (i + 1))
        for (let j = 0; j < col.length + 2; j++) {
            let tabCell = tr.insertCell(-1)
            tabCell.innerHTML = arrShopping[i][col[j]]
            if (j === 3) {
                tabCell.innerHTML = 'Delete'
                tabCell.setAttribute('type', 'button')
                tabCell.setAttribute('value', 'Delete Row')
                tabCell.setAttribute('class', 'delButton')
                tabCell.setAttribute('onclick', 'deleteRow(this.id)')
                tabCell.setAttribute('id', 'd' + (arrShopping[i][col[0]]))
            }
            if (j === 4) {
                tabCell.innerHTML = 'Update'
                tabCell.setAttribute('type', 'button')
                tabCell.setAttribute('value', 'Update Row')
                tabCell.setAttribute('class', 'updtButton')
                tabCell.setAttribute('onclick', 'rediToUpdt(this.id)')
                tabCell.setAttribute('id', 'u' + (arrShopping[i][col[0]]))
            }
        }
    }
    let divContainer = document.getElementById('showTable')
    divContainer.innerHTML = ""
    divContainer.appendChild(table)
    let divContainerBt = document.getElementById('showBtn')
    divContainerBt.innerHTML = ""
}
//This tells the backend to delete a row at given ID
function deleteRow(id) {
    if (window.confirm('Are you sure you want to delete this item?')) {
    id = trim(id)
    id = parseInt(id)
    let data = {
        Id: id
    }
    fetch("/delete_row", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data)
            console.log(result)
        })
    }).catch((error) => {
        console.log(error)
    })
    setTimeout(() => {window.location.reload()}, 350)
}
}
//This trims the first char in given string, used for trimming ID
function trim(s) {
    let sTrimmed = s.substring(1)
    return sTrimmed
}
//This redirects to update.html and also stores the given ID in the URL
function rediToUpdt(id) {
    id = trim(id)
    url = 'http://127.0.0.1:5500/update.html?id=' + encodeURIComponent(id)
    document.location.href = url
}

