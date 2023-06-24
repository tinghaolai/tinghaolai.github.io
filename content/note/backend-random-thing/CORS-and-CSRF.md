---
title: "CORS-and-CSRF.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["backend-random-thing"]
---



---

# Noun

CORS: Cross-Origin Resource Sharing

CSRF: Cross-Site Request Forgery

Same Origin Policy

front-end and back-end separation

## CORS
Even though browsers enforce the same-origin policy, 
preventing the sending of cookies across different domains, 
it is still possible for unintended actions to occur if the backend API does not have any checks in place. 
Therefore, it is necessary to check if the requests are from the same origin. 

In the case of a frontend-backend separation, 
if the host is on a different domain, 
the frontend domain needs to be included in the allow list.

## CSRF
Why do we still need CSRF tokens despite having CORS protection?
Firstly,

CORS only protects against cookie theft. 
If a browser does not support CORS, 
cookie theft can still occur.
Even with CORS protection on the backend, 
the source can be forged, including the referrer.
As mentioned earlier, if the backend does not check cookies, 
unintended actions can still occur and result in damage.

### How does CSRF token work?
The server sends a token to the client.
The client includes this token in the submitted form.
If the token is invalid, the server rejects the request.

### How does Laravel handle CSRF protection?
The CSRF token is rendered on the frontend page and then checked against the session on the backend.

### How can CSRF tokens be protected from being stolen?
It is generally not a major concern. 
If it is stolen, 
it is likely that not only the token but all information has been compromised.

### How to handle CSRF protection in a frontend-backend separation?
The backend encrypts the token and stores it on the frontend. 
When making a request, 
the token is sent back to the backend. 
Due to the browser's same-origin policy, 
the token should not be accessible even in the event of a CSRF attack. 

## Summarize:

If CSRF tokens are implemented, is CORS not necessary?

Although it is recommended to implement both for comprehensive defense, 
theoretically, yes. 
CSRF tokens verify the trusted source,
assuming that premise, CORS 
is not needed. However, CORS still has two risks:

The source can be forged.
Browsers that do not support CORS can still steal cookies.
The backend does not check cookies.
Conclusion:
With proper API design, normal browsers, and framework handling, there is little need to worry about this issue. 
However, for analysis purposes:

* CORS does not guarantee complete security.
* CSRF tokens do not guarantee complete security either, unless the user is using a very poor browser or has installed malicious plugins.
  * SSR token: Generally the most secure option.
  * Frontend-backend separation token: Secure depending on the browser.

---

