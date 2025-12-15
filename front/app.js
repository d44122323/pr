
const navLinks = document.querySelectorAll('.nav-menu a');
const btnRecord = document.querySelector('.btn-record');


function scrollToBlock(selector) {
    const block = document.querySelector(selector);
    if (block) {
        block.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }
}

navLinks.forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault(); 
        const text = link.textContent.trim();

        switch(text) {
            case 'Главная':
                scrollToBlock('#hero');
                break;
            case 'Услуги':
                scrollToBlock('#services');
                break;
            case 'Специалисты':
                scrollToBlock('#specialists');
                break;
            case 'Контакты':
                scrollToBlock('#contacts');
                break;
        }
    });
});


btnRecord.addEventListener('click', () => {
    scrollToBlock('#contacts');
});


const serviceButtons = document.querySelectorAll('.btn-service');

const specialistButtons = document.querySelectorAll('.btn-specialist');


function goToContacts() {
    const block = document.querySelector('#contacts');
    if (block) {
        block.scrollIntoView({ behavior: 'auto', block: 'start' });
    }
}


serviceButtons.forEach(btn => {
    btn.addEventListener('click', goToContacts);
});


specialistButtons.forEach(btn => {
    btn.addEventListener('click', goToContacts);
});


const btnHero = document.querySelector('.btn-hero');

btnHero.addEventListener('click', () => {
    const contacts = document.querySelector('#contacts');
    if (contacts) {
        window.location.href = '#contacts';
    }
});



document.addEventListener("DOMContentLoaded", () => {
    const recordButtons = document.querySelectorAll(".btn-record, .btn-hero, .btn-specialist");
    const contactForm = document.getElementById("contacts");

    recordButtons.forEach(btn => {
        btn.addEventListener("click", () => {
            contactForm.scrollIntoView({ behavior: "smooth" });
        });
    });
});








