let datas = []

function getData(event) {
    event.preventDefault()

    let title = document.getElementById("input-blog-title").value
    let content = document.getElementById("input-blog-content").value
    let image = document.getElementById("input-blog-image").files[0]

    image = URL.createObjectURL(image[0])

    let data = {
        title,
        content,
        image,
        date: new Date(),
        author: "Muhammad Risky Fadhiila"
    }

    datas.push(data)
    console.log(datas)
    showData()
}

const showData = () => {
    document.getElementById("contents").innerHTML = ''
    for (let i = 0; i < datas.length; i++) {
        document.getElementById("contents").innerHTML += `
        <div class="blog-list-item">
          <div class="blog-image">
            <img src="${datas[i].image}" alt="" />
          </div>
          <div class="blog-content">
            <div class="btn-group">
              <button class="btn-edit">Edit Post</button>
              <button class="btn-post">Post Blog</button>
            </div>
            <h1>
              <a href="blog-detail.html" target="_blank"
                >${datas[i].title}</a
              >
            </h1>
            <div class="detail-blog-content">
              
            ${createTime(time)}${datas[i].date} | ${datas[i].author}
            </div>
            <p>
              ${datas[i].content}
            </p>

            <div style="float:right; margin: 10px">
              <p style="font-size: 15px; color:grey">1 minutes ago</p>
            </div>

          </div>
        </div>`
    }
}

function createTime(time) {
    let years = time.getFullYear()
    let month = time.getMonth()
    let date = time.getDate()
    let hours = time.getHours()
    let minutes = time.getMinutes()
    return (`${date} ${month} ${years} ${hours} : ${minutes} WIB`)
}