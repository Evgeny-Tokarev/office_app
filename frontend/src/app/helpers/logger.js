/* eslint-disable no-console */

class Logger {
    constructor () {
        // this.isDebug = process.env.NODE_ENV !== 'production'
        this.isDebug = false
    }

    isDebugMode () {
        return this.isDebug
    }

    groupCollapsed (...data) {
        if (this.isDebug) console.groupCollapsed(...data)
    }

    group (...data) {
        if (this.isDebug) console.group(...data)
    }

    groupEnd () {
        if (this.isDebug) console.groupEnd()
    }

    warn (...data) {
        if (this.isDebug) console.warn(...data)
    }

    error (...data) {
        if (this.isDebug) console.error(...data)
    }

    log (...data) {
        if (this.isDebug) console.log(...data)
    }
}

export default new Logger()
