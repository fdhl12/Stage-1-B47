// class Testimonial {
//     constructor(quote, image) {
//         this._quote = quote
//         this._image = image
//     }

//     get quote() {
//         return this._quote
//     }

//     get image() {
//         return this._image
//     }

//     get html() {
//         return `<div class="testimonial">
//             <img src="${this.image}" class="profile-testimonial" />
//             <p class="quote">"${this.quote}"</p>
//             <p class="author">- ${this.author}</p>
//         </div>`
//     }
// }

// class AuthorTestimonial extends Testimonial {
//     constructor(author, quote, image) {
//         super(quote, image)
//         this._author = author
//     }

//     get author() {
//         return this._author
//     }
// }

// class CompanyTestimonial extends Testimonial {
//     constructor(company, quote, image) {
//         super(quote, image)
//         this._company = company
//     }

//     get author() {
//         return this._company + " Company"
//     }
// }

// const testimonial1 = new AuthorTestimonial("Fadil", "Mantap, keren banget jasanya", "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60")

// const testimonial2 = new AuthorTestimonial("Surya ", "oke sih, oke lah", "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60")

// const testimonial3 = new CompanyTestimonial("Budi", "Gege gaming!", "https://images.unsplash.com/photo-1615109398623-88346a601842?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTd8fG1hbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60")

// let testimonialData = [testimonial1, testimonial2, testimonial3]
// let testimonialHTML = "";

// for (let i = 0; i < testimonialData.length; i++) {
//     testimonialHTML += testimonialData[i].html
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML

const testimonialsData = [{
        author: "Fadil",
        quote: "Goodjobs!!",
        image: "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
        rating: 5,
    },
    {
        author: "Surya",
        quote: "Don't Pesimis",
        image: "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
        rating: 4,
    },
    {
        author: "Budi",
        quote: "Semangat",
        image: "https://images.unsplash.com/photo-1615109398623-88346a601842?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTd8fG1hbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
        rating: 4,
    }
];

function allTestimonials() {
    let testimonialHTML = "";

    testimonialsData.forEach(function (item) {
        testimonialHTML += `<div class="testimonial">
                                    <img
                                    src="${item.image}" alt=""/>
                                <p class="quote">${item.quote}</p>
                                <p class="author">${item.author}</p>
                                <p class="author">${item.rating}<i class="fa-solid fa-star"></i></p>
                            </div>`
    });
    document.getElementById("testimonials").innerHTML = testimonialHTML;

}

allTestimonials();

function filterTestimonials(rating) {
    let testimonialHTML = "";

    const testimonialsFiltered = testimonialsData.filter(function (item) {
        return item.rating === rating;
    });

    if (testimonialsFiltered.length === 0) {
        testimonialHTML += `<h1>Data not found!</h1>`;
    } else {
        testimonialsFiltered.forEach(function (item) {
            testimonialHTML += `<div class="testimonial">
                                    <img
                                        src="${item.image}"
                                        class="profile"
                                    />
                                    <p class="quote">${item.quote}</p>
                                    <p class="author">- ${item.author}</p>
                                    <p class="rating">${item.rating} <i class="fa-solid fa-star"></i></p>
                                </div>
                            `;
        });
    }

    document.getElementById("testimonials").innerHTML = testimonialHTML;
}