function toggleShowNav() {
    const navSm = document.getElementById("nav-sm")
    navSm.classList


    if ([...navSm.classList].includes("d-none"))
        navSm.classList = "d-show transition"
    else navSm.classList = "d-none transition"
}