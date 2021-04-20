/* tslint:disable */
/* eslint-disable */
// Code is generated by github.com/swaggest/swac <version>, DO NOT EDIT. 🤖

/**
 * Swagger Petstore
 * Version: 1.0.0
 * A sample API that uses a petstore as an example to demonstrate features in
 * the OpenAPI 3.0 specification
 * @constructor
 * @param {string} baseURL - Base URL.
 */
function APIClient(baseURL) {
    // Trim trailing backslash.
    this.baseURL = (baseURL.charAt(baseURL.length - 1) === '/') ? baseURL.slice(0, -1) : baseURL;
    /** @type {?PrepareRequest} - Callback to prepare request before sending. */
    this.prepareRequest = null;
}

/**
 * @callback PrepareRequest
 * @param {XMLHttpRequest} value
 */

/**
 * Returns all pets from the system that the user has access to
 * Nam sed condimentum est. Maecenas tempor sagittis sapien, nec rhoncus sem
 * sagittis sit amet. Aenean at gravida augue, ac iaculis sem. Curabitur odio
 * lorem, ornare eget elementum nec, cursus id lectus. Duis mi turpis,
 * pulvinar ac eros ac, tincidunt varius justo. In hac habitasse platea
 * dictumst. Integer at adipiscing ante, a sagittis ligula. Aenean pharetra
 * tempor ante molestie imperdiet. Vivamus id aliquam diam. Cras quis velit
 * non tortor eleifend sagittis. Praesent at enim pharetra urna volutpat
 * venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent enim
 * enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor,
 * consectetur non augue vel, vehicula interdum libero. Morbi euismod sagittis
 * libero sed lacinia.

 * Sed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu
 * condimentum ligula luctus nec. Phasellus semper velit eget aliquet
 * faucibus. In a mattis elit. Phasellus vel urna viverra, condimentum lorem
 * id, rhoncus nibh. Ut pellentesque posuere elementum. Sed a varius odio.
 * Morbi rhoncus ligula libero, vel eleifend nunc tristique vitae. Fusce et
 * sem dui. Aenean nec scelerisque tortor. Fusce malesuada accumsan magna vel
 * tempus. Quisque mollis felis eu dolor tristique, sit amet auctor felis
 * gravida. Sed libero lorem, molestie sed nisl in, accumsan tempor nisi.
 * Fusce sollicitudin massa ut lacinia mattis. Sed vel eleifend lorem.
 * Pellentesque vitae felis pretium, pulvinar elit eu, euismod sapien.
 * @param {FindPetsRequest} req - request parameters.
 * @param {ArrayNewPetPetAllOf1Callback} onOK
 */
APIClient.prototype.findPets = function (req, onOK) {
    var x = new XMLHttpRequest();
    x.onreadystatechange = function () {
        if (x.readyState !== XMLHttpRequest.DONE) {
            return
        }
    
        switch (x.status) {
            case 200:
                if (typeof(onOK) == 'function') {
                    onOK(JSON.parse(x.responseText));
                }
                break;
            default:
                throw {err: 'unexpected response', data: x}
        }
    };
    
    var url = this.baseURL + '/pets?';
    if (req.tags != null) {
        url += 'tags=' + encodeURIComponent(req.tags) + '&'
    }
    if (req.limit != null) {
        url += 'limit=' + encodeURIComponent(req.limit) + '&'
    }
    url = url.slice(0, -1)

    x.open("GET", url, true);
    if (typeof(this.prepareRequest) == 'function') {
        this.prepareRequest(x);
    }
    
    x.send();
}

/**
 * Creates a new pet in the store.  Duplicates are allowed
 * @param {PostPetsRequest} req - request parameters.
 * @param {NewPetPetAllOf1Callback} onOK
 */
APIClient.prototype.postPets = function (req, onOK) {
    var x = new XMLHttpRequest();
    x.onreadystatechange = function () {
        if (x.readyState !== XMLHttpRequest.DONE) {
            return
        }
    
        switch (x.status) {
            case 200:
                if (typeof(onOK) == 'function') {
                    onOK(JSON.parse(x.responseText));
                }
                break;
            default:
                throw {err: 'unexpected response', data: x}
        }
    };
    
    var url = this.baseURL + '/pets?';
    url = url.slice(0, -1)

    x.open("POST", url, true);
    if (typeof(this.prepareRequest) == 'function') {
        this.prepareRequest(x);
    }
    if (typeof req.body !== 'undefined') {
        x.setRequestHeader("Content-Type", "application/json; charset=utf-8");
        x.send(JSON.stringify(req.body))
        return
    }

    x.send();
}

/**
 * Returns a user based on a single ID, if the user does not have access to
 * the pet
 * @param {GetPetsIdRequest} req - request parameters.
 * @param {NewPetPetAllOf1Callback} onOK
 */
APIClient.prototype.getPetsId = function (req, onOK) {
    var x = new XMLHttpRequest();
    x.onreadystatechange = function () {
        if (x.readyState !== XMLHttpRequest.DONE) {
            return
        }
    
        switch (x.status) {
            case 200:
                if (typeof(onOK) == 'function') {
                    onOK(JSON.parse(x.responseText));
                }
                break;
            default:
                throw {err: 'unexpected response', data: x}
        }
    };
    
    var url = this.baseURL + '/pets/' + encodeURIComponent(req.id) +
    '?';
    url = url.slice(0, -1)

    x.open("GET", url, true);
    if (typeof(this.prepareRequest) == 'function') {
        this.prepareRequest(x);
    }
    
    x.send();
}

/**
 * deletes a single pet based on the ID supplied
 * @param {DeletePetsIdRequest} req - request parameters.
 * @param {RawCallback} onNoContent
 */
APIClient.prototype.deletePetsId = function (req, onNoContent) {
    var x = new XMLHttpRequest();
    x.onreadystatechange = function () {
        if (x.readyState !== XMLHttpRequest.DONE) {
            return
        }
    
        switch (x.status) {
            case 204:
                if (typeof(onNoContent) == 'function') {
                    onNoContent(x);
                }
                break;
            default:
                throw {err: 'unexpected response', data: x}
        }
    };
    
    var url = this.baseURL + '/pets/' + encodeURIComponent(req.id) +
    '?';
    url = url.slice(0, -1)

    x.open("DELETE", url, true);
    if (typeof(this.prepareRequest) == 'function') {
        this.prepareRequest(x);
    }
    
    x.send();
}
