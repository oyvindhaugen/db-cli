function redirectToInsert() {
    alert("Hello")
}

let oXHR = new XMLHttpRequest()
oXHR.onreadystatechange = reportStatus
oXHR.open("GET", "../selectQuery.json", true)
oXHR.send()
function reportStatus() {
    if (oXHR.readyState === 4) {
        createTableFromJSON(this.responseText)
    }
}
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

    for (let i = 0; i < col.length + 1; i++) {
        let th = document.createElement("th")
        th.innerHTML = col[i]
        tr.appendChild(th)
        if (i === 3) {
            th.innerHTML = 'delBtn'
            tr.appendChild(th)
        }
    }


    for (let i = 0; i < arrShopping.length; i++) {
        let bt = addAttrToButton(i)
        tr = table.insertRow(-1)
        tr.setAttribute('id', "tr" + (i + 1))
        for (let j = 0; j < col.length + 1; j++) {
            let tabCell = tr.insertCell(-1)
            tabCell.innerHTML = arrShopping[i][col[j]]
            if (j == 3) {
                tabCell.innerHTML = 'Delete'
                tabCell.setAttribute('type', 'button')
                tabCell.setAttribute('value', 'Delete row')
                tabCell.setAttribute('class', 'delButton')
                tabCell.setAttribute('onclick', '')
                tabCell.setAttribute('id', 'b' + (i + 1))
                tabCell.addEventListener('click', function () {
                    console.log('test', tabCell.id)
                    //now it logs to console its id
                    //use this to tell the database to delete a certain item
                    //maybe an alert to confirm
                })
            }
        }
    }

    let divContainer = document.getElementById('showTable')
    divContainer.innerHTML = ""
    divContainer.appendChild(table)
    let divContainerBt = document.getElementById('showBt')
    divContainerBt.innerHTML = ""
}

function addAttrToButton(id) {
    let input = document.createElement('input')
    input.setAttribute('type', 'button')
    input.setAttribute('value', 'Delete row')
    input.setAttribute('class', 'delButton')
    input.setAttribute('onclick', 'test()')
    input.setAttribute('id', 'b' + (id + 1))
    console.log(input)
    return input
}

function deleteRow(id) {
    let data = {
        Id: 3
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
}
// function test() {
//     let i = 1
//     let current = document.getElementById("b" + i)
//     console.log(current.id)
// }
//This function trims away the first character in the id of button
//so it can be used to delete from table
function trim(s) {
    let sTrimmed = s.substring(1)
    return sTrimmed
}
