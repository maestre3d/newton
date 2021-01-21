import React from 'react'
import { Fragment } from 'react'
import { Helmet } from 'react-helmet'
import { Input } from 'baseui/input'
import { APPLICATION_DOMAIN, APPLICATION_NAME } from '../../../internal/shared/domain/newton'

function Home() {
    const [value, setValue] = React.useState('')
    return (
        <Fragment>
            <Helmet>
                <meta charSet='utf-8' />
                <title>{`${APPLICATION_NAME} - Home`}</title>
                <link rel='canonical' href={`https://${APPLICATION_DOMAIN}/`} />
            </Helmet>
            <div className='p-8'>
                <Input
                    value={value}
                    onChange={event => setValue(event.currentTarget.value)}
                    placeholder="Controlled Input"
                />
            </div>
        </Fragment>
    )
}

export default Home