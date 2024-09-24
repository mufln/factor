const form = document.querySelector('.login');
const enterButton = document.querySelector('#enter');

enterButton.addEventListener('click', (e) => {
  e.preventDefault();

  const nickname = document.querySelector('#name').value;
  const password = document.querySelector('#password').value;
  const repPassword = document.querySelector('#rep-password').value;

  if (password !== repPassword) {
    alert('Пароли не совпадают');
    return;
  }

  const formData = new FormData();
  formData.append('nickname', nickname);
  formData.append('password', password);

  fetch('/register', {
    method: 'POST',
    body: formData
  })
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error(error));
});