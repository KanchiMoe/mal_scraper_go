package src

const ROOT_URL string = "https://myanimelist.net/"

const PEOPLE_URL string = "people/"          // https://myanimelist.net/people/000/
const CHARACTER_URL string = "character/"    // https://myanimelist.net/character/000/
const COMPANY_URL string = "anime/producer/" // https://myanimelist.net/anime/producer/000
const GENRE_URL string = "anime/genre/"      // https://myanimelist.net/anime/genre/000

// users
const PROFILE_URL string = "profile/" // https://myanimelist.net/profile/Name

// thingy urls
const REPORT_SLUG string = "modules.php?go=report&type=profile&id=" // root + this + user id
const COMMENTS_SLUG string = "comments.php?id="                     // root + this + userid

// regex
// Number in brackets. Example: (123)
const REGEX_INT_IN_BRACKETS string = `\((\d+)\)`
