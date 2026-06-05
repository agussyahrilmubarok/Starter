function handleEvent(eventType, selector, handler) {
    document.addEventListener(eventType, function (event) {
        if (event.target.matches(selector + ', ' + selector + ' *')) {
            handler.apply(event.target.closest(selector), arguments);
        }
    });
}

handleEvent('submit', '.js-submit-confirm', function (event) {
    if (!confirm(this.getAttribute('data-confirm-message'))) {
        event.preventDefault();
        return false;
    }
    return true;
});