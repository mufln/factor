var script = document.createElement('script');
script.src = 'https://code.jquery.com/jquery-3.7.1.min.js'; // Check https://jquery.com/ for the current version
document.getElementsByTagName('head')[0].appendChild(script);

class ChatInfo {
    constructor(group_pic_path, name, id)
    {
        this.group_pic_path = group_pic_path;
        this.name = name;
        this.id = id;
    }
    getName()
    {
        return this.name;
    }
}

var chatsBTN = document.getElementById("show-chats");
var list = document.getElementById("chats");
var shown = 0;
var menu = document.getElementById("menu");

var chatInfo = new Array('Chat1', 'BlaBlaChat2', 'Trututu', 'Chat3','Chat4', 'SOSChatik');
var chatiki = "";
for (i = 0; i < chatInfo.length; i++)
{
    chatiki = chatiki + `
    <div class = "chat" id="chat">
        <div class="user-info">
            <div class="avatar"><div class = "firstLetter"><div class = "letter">${Array.from(chatInfo[i])[0]}</div></div></div>
            <div class="text-info">
                <div class="name" id = "chat-name">${chatInfo[i]}</div> 
            </div>
        </div>
    </div>
    `
}
document.getElementById("chats").innerHTML = chatiki;

chatsBTN.onclick = function()
{   
    if (shown == 0)
    {   
        list.style.display = "block";
        menu.style.display = "none";
        chatsBTN.style.transform = "rotate(180deg)";
        shown = 1;
    }
    else
    {
        list.style.display = "none";
        menu.style.display = "block";
        chatsBTN.style.transform = "rotate(0deg)";
        shown = 0;
    }
}

var element = document.getElementsByClassName("chat");
for (i = 0;  i < element.length; i++){
    element[i].addEventListener('click', ShowMessages, false);
}

function ShowMessages(){
    
    document.getElementById("message-column").style.backgroundColor = "white";
    document.getElementById("message-column").innerHTML = `
    <header style = "padding-right: 1%; box-shadow: inset 0 -40px 0 0 #0000000D, 0px 10px 20px 0px #0000000D; " >
          <div class="chat-info">
              <div class="avatar"><div class="avatar"><div class = "firstLetter"><div class = "letter">C</div></div></div></div>
              <div class="name" >Chat</div>
          </div>
      </header>
  
          <div class="messages" id="messages" style="overflow-y:scroll; overflow-x:clip;">
          </div>
  
      <div class="footer">
          <form class="footer-inner" id="sender-form">
          <div class="text-field" contenteditable="true">
              <textarea method="POST" placeholder="Write a message..." rows="1" oninput="auto_grow(this)"></textarea>
          </div>
          <div class="pin">
          </div>
          <div class="send">
              <label for="sender">
                  <svg width="40" height="40" viewBox="0 0 44 44" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M42 2L20 24M42 2L28 42L20 24M42 2L2 16L20 24" stroke="#01A274" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
                  </svg>
              </label>
              <input id="sender" type="submit"/>
          </div>
      </form>
      </div>
  </div>`;
      var messagesInfo = new Array('Bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla bla',
          'aaaaaaaaaaaaaaaaaaahahahhahahaaaaaaaaaaaaaaaaaaaaaaaaaaaahHAHAHAHAHAHAHAHAHHAHAHAHHAHAHHAHHA', 'yeah');
      var messages = "";    
      for (i = 0; i < 7; i++){    
          messages = messages + `
          <div class="message">
                  <div class="avatar" style="width: 30px; height: 30px;"><img src="https://img.freepik.com/free-photo/adorable-kitty-with-monochrome-wall-her_23-2148955138.jpg?t=st=1725978880~exp=1725982480~hmac=7aac8bb70992b23e4796d28ae5bc0117d552a81c58591434a63a17ffa9119ae9&w=1380"
                                                                              style="width: 30px; height: 30px;"></div>
                  <div class="text">
                      <p style="margin: 10px 0px 0px;">${messagesInfo[i%3]}</p>
                      <div class = "time-group"> 
                          <div class="time">10.09.2024</div>
                          <div class="time">20:30</div>
                      </div>
                  </div>
              </div>`
      }        
  //// добавляем сообщения
      document.getElementById("messages").innerHTML = messages;
  
  //// убираем выделение выделенного ранее чата
      var element = document.getElementsByClassName("chat-selected");
      if (element[0] != null){
          element[0].getElementsByClassName("name-selected")[0].className = "name";
          element[0].className = "chat";
      }
  
  ////выделяем новый чат
      this.className = "chat-selected";
      this.getElementsByClassName("name")[0].className = "name-selected";
  
      var backButton = document.getElementById("back");
  
      window.onresize = function (){
          if (window.innerWidth < 720 && backButton!=null){
              if (document.getElementById("message-column").style.backgroundColor == "white"){
                  backButton.innerHTML = `<div class="avatar"><img src="back.png"></div>`;
              }
              backButton.addEventListener('click', ShowChats, false);
          }
          else if (backButton!=null){
              backButton.innerHTML = ``;
              document.getElementById("changeview").innerHTML = `
              @media (max-width: 720px) {
                  .chats-column {display: none;}
              }`
          }
      }
  
      /////////////здесь мы автоматом крутим вниз
      var scrollDiv = document.getElementById("messages");
      if (scrollDiv != null) {
          scrollDiv.scrollTo(0, scrollDiv.scrollHeight);
      } 
  
      document.getElementById("changeview").innerHTML = `
      @media (max-width: 720px) {
          .chats-column {display: none;}
      }`
  
      if (window.innerWidth < 720 && backButton!=null){
          if (document.getElementById("message-column").style.backgroundColor == "white"){
              backButton.innerHTML = `<div class="avatar"><img src="back.png"></div>`;
          }
          backButton.addEventListener('click', ShowChats, false);
      }
      else if (backButton!=null){
          backButton.innerHTML = ``;
      }
  
      var messages = document.getElementById("messages");
      messages.addEventListener('scroll', ShowScroll, false);
}