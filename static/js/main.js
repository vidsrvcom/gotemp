// Main JavaScript file
document.addEventListener('DOMContentLoaded', function() {
    console.log('Go Web Template loaded successfully!');
    
    // Add click handlers to endpoint rows for API testing
    document.querySelectorAll('.endpoint').forEach(function(endpoint) {
        endpoint.style.cursor = 'pointer';
        endpoint.addEventListener('click', function() {
            var path = this.querySelector('.path').textContent;
            var method = this.querySelector('.method').textContent;
            console.log('API Endpoint: ' + method + ' ' + path);
        });
    });
});
