//funtion
function getData() {
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let address = document.getElementById("address").value
    //coditional
    if (name == "") {
        return alert("nama harus diisi")
    } else if (email == "") {
        return alert("email harus diisi")
    } else if (phone == "") {
        return alert("email harus diisi")
    } else if (email == "") {
        return alert("email harus diisi")
    } else if (email == "") {
        return alert("email harus diisi")
    }

    const emailReceiver = "riskyfadhiila52@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}? subject= ${subject}&body=Hello Nama Saya ${name},
    alamat saya ${address}, bisakah anda menghubungi saya di ${phone}`
    a.click()
    //data
    let data = {
        name,
        email,
        phone,
        subject,
        address
    }
}