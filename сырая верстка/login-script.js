const form = document.querySelector('.login');
const enterButton = document.querySelector('#enter');

enterButton.addEventListener('click', (e) => {
  e.preventDefault();

  const nickname = document.querySelector('#name').value;
  const password = document.querySelector('#password').value;

  const formData = new FormData();
  formData.append('nickname', nickname);
  formData.append('password', password);

  fetch('/login', {
    method: 'POST',
    body: formData
  })
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error(error));

  const template = Handlebars.compile($("#index").html());
  const html = template(data);
  const newPage = document.createElement("div");
  newPage.innerHTML = html;
  document.body.appendChild(newPage);
});

