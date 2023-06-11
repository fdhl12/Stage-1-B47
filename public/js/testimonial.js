const testimonialPromise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open('GET', 'https://api.npoint.io/fa58c458e0eca036cc96', true)

    xhr.onload = function () {
        // 200 : ok
        if (xhr.status === 200) {
            resolve(JSON.parse(xhr.response))
        } else {
            reject('Error loading data!')
        }
    }
    xhr.onerror = function () {
        reject("Network Error!")
    }
    xhr.send()
});
// (async () => console.log(await testimonialPromise))()

let rating = []

async function getAllTestimonials() {
    const response = await testimonialPromise;
    //   console.log(response);

    let testimonialHTML = "";
    response.forEach(function (item) {
        testimonialHTML += `<div class="testimonial">
                                <img src="${item.image}" alt="" class="profile-testimonial"/>
                                <p class="coment">${item.quote}</p>
                                <p class="author">- ${item.author}</p>
                                <p class="author"> ${item.rating}<i class="fa-solid fa-star"></i></p>
                            </div>`;
    });

    document.getElementById("testimonials").innerHTML = testimonialHTML;
}

getAllTestimonials();

const ratingStars = [...document.getElementsByClassName("rating__star")]
//spread operator berfungsi untuk mengeluarkan element didalam array

async function filterTestimonials(rating) {
    const ratingTestimonial = (rating) => {
        let tempRating = []
        for (let i = 1; i <= rating; i++) {
            tempRating = [...tempRating, `<i class="rating__star fas fa-star"></i>`]
        }
        return tempRating
        // console.log("rating", tempRating);
    }
    let testimonialHTML = '';
    // console.log(!!1)
    let testimonialFiltered = await testimonialPromise //nilai default
    if (rating) //jika pemanggilan fungsi filter testimoninal menggunakan rating
        testimonialFiltered = (await testimonialPromise).filter((item) => { //ubah nilai testimonial filtered dengan data ratingnya sama
            return item.rating === rating
        })

    //ternary operator
    // const testimonialFiltered = rating ? testimonialData.filter((item) => {
    //     return item.rating === rating
    // }) : testimonialData

    if (!rating) {
        ratingStars.map((item) =>
            item.classList = "rating__star far fa-star"
        )
    }
    if (!testimonialFiltered.length) {
        testimonialHTML = `<h1> Data not found! </h1>`
    } else {
        testimonialFiltered.forEach((item) => {
            // `template literal` << backtik
            testimonialHTML += `<div class="testimonial">
                <img src="${item.image}" class="profile-testimonial" />
                <p class="quote">"${item.quote}"</p>
                <p class="author">- ${item.author}</p>
               <p class="author"> ${ratingTestimonial(item.rating).map(function (e) {
                return e 
              }).join('')}</p>
            </div>`
        })
    }

    document.getElementById('testimonials').innerHTML = testimonialHTML
}
filterTestimonials()

// function executeRating(stars) {
//     const starClassActive = "rating__star fas fa-star";
//     const starClassInactive = "rating__star far fa-star";
//     const starsLength = stars.length
//     let i
//     // i atau inisiasi
//     stars.map((star) => {
//         star.onclick = () => {
//             i = stars.indexOf(star)
//             console.log(i);

//             if (star.className === starClassInactive) {
//                 for (i; i >= 0; --i) stars[i].className = starClassActive;
//             } else {
//                 for (i; i < starsLength; ++i)
//                     if (!!stars[i + 1]) stars[i + 1].className = starClassInactive;
//             }
//             filterTestimonials(stars.indexOf(star) + 1)
//         }
//     })
// }
// executeRating(ratingStars);