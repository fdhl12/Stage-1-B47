let datas = []
let dataFull = []

function getData(event) {
    event.preventDefault()
    let projectName = document.getElementById("projectName").value
    let startDate = document.getElementById("startDate").value
    let endDate = document.getElementById("endDate").value
    let Description = document.getElementById("Description").value
    let image = document.getElementById("image").files[0]

    image = URL.createObjectURL(image)

    let checkboxs = document.querySelectorAll('input[type="checkbox"]:checked');
    checkboxs.forEach((checkboxs) => {
        datas.push(checkboxs.value)
    })
    let tech = {
        projectName,
        startDate,
        endDate,
        Description,
        tech: datas,
        image,
    }

    dataFull.push(tech)
    datas = []
    showData()
}
const showData = () => {
    document.getElementById("contents").innerHTML = ``;
    for (let i = 0; i < dataFull.length; i++) {
        let validasi = dataFull[i].tech;
        document.getElementById("contents").innerHTML += `
       <div class="card">
  <div class="upper">
      <img src="${dataFull[i].image}" class="object-cover" alt="">
  </div>
  <div>
      <div class="title-card"> <a href="detail-project.html?projectName=${
        dataFull[i].projectName
      }&endDate=${dataFull[i].endDate}&startDate=${
      dataFull[i].startDate
      }&image=${dataFull[i].image}&description=${dataFull[i].Description}&tech=${
      validasi[i]
      }&duration=${getDuration(dataFull[i].startDate, dataFull[i].endDate)}"> ${
      dataFull[i].projectName
      } - ${new Date(dataFull[i].endDate).getFullYear()} </a></div>
      <div class="bulan"> durasi: ${getDuration(
        dataFull[i].startDate,
        dataFull[i].endDate
      )} </div>
  </div>
      <div>
          <p class="text-card">${dataFull[i].Description}</p>
      </div>
  <div>
      <div class="logo">
          ${validasi.map(function (e) {
              return '<i class="' + e + '"></i>'
            }).join('')}
      </div>
      <div class="action">
          <button>edit</button>
          <button>delete</button>
      </div>
  </div>
</div> `
    }
}

const getDuration = (timeStart, timeEnd) => {
    let distance = new Date(timeEnd) - new Date(timeStart)
    const miliSecond = 1000;
    const yearDistance = Math.floor(
        distance / (12 * 4 * 7 * 24 * 60 * 60 * miliSecond)
    )
    if (yearDistance > 0) {
        return yearDistance + " Years Ago";
    } else {
        const monthDistance = Math.floor(
            distance / (4 * 7 * 24 * 60 * 60 * miliSecond)
        )
        if (monthDistance > 0) {
            return monthDistance + " Month Now";
        } else {
            const weekDistance = Math.floor(
                distance / (7 * 24 * 60 * 60 * miliSecond)
            )
            if (weekDistance > 0) {
                return weekDistance + " Week Now";
            } else {
                const dayDistance = Math.floor(distance / (24 * 60 * 60 * miliSecond))
                if (dayDistance > 0) {
                    return dayDistance + " Day Now";
                } else {
                    const closeDistance = Math.floor(distance / (60 * 60 * miliSecond))
                    if (closeDistance > 0) {
                        return "Just Now"
                    } else {
                        if (closeDistance < 0) {
                            return alert(
                                "Tanggal tidak valid : Start Date harus > dari End Date"
                            )
                        }
                    }
                }
            }
        }
    }
}