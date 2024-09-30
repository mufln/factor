function generateRandomShadeOfColor(baseColor) {
    // Convert the base color to RGB
    const baseRgb = hexToRgb(baseColor);
  
    // Generate a random factor to adjust the brightness of the color
    const factor = Math.random() * 0.5 + 0.5; // adjust brightness by 50% to 150%
  
    // Calculate the new RGB values
    const newR = Math.round(baseRgb.r * factor);
    const newG = Math.round(baseRgb.g * factor);
    const newB = Math.round(baseRgb.b * factor);
  
    // Return the new color as a hex string
    return `rgb(${newR}, ${newG}, ${newB})`;
  
}

function hexToRgb(hex) {
    const r = parseInt(hex.substring(1, 3), 16);
    const g = parseInt(hex.substring(3, 5), 16);
    const b = parseInt(hex.substring(5, 7), 16);
    return { r, g, b };
}
/*
const vueApp = Vue.createApp({
    el: '#app',
    setup() {
        const title = 'Фактор';
        return { title }
    },
    mounted() {
        $('#app').load('index.html');
    },
    methods: {
    },
});
vueApp.mount('#app');
*/
$('#app').load('index.html');

$("#tasks-button").on('click', (function() {
    $('#content').load('task.html');
}))

$("#teams-button").on('click', (function() {
    var clr;
    var teams = "";
    for (var i = 0; i < 10; i++){
        clr = generateRandomShadeOfColor("#01A274");
        teams += `<li class="team-card">
                <div class="team-avatar"  style="background-color:${clr}">
                </div>
                <span class="team-name">Мемный отдел sghjsbcmzbcmnzbcnmzbnmbzc</span>    
            </li>`;
    };
    $('#app').load('teams.html');
}))

$("#new-team").on('click', (function() {
    $('#app').load('task.html');
}))