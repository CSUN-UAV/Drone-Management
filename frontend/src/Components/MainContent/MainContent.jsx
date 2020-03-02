import React from 'react';

import { makeStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Divider from '@material-ui/core/Divider';

import { useEffect, useRef } from 'react'


const useStyles = makeStyles(theme => ({
    root: {
      display: 'flex',
    },
    toolbar: theme.mixins.toolbar,
    content: {
      flexGrow: 1,
      backgroundColor: theme.palette.background.default,
      padding: theme.spacing(3),
    },
    heading: {
      width: '100%',
      maxWidth: 500,
      marginTop: '2vh',
      marginBottom: '2vh',
      color: '#D22030',
      fontWeight: 'bold'
    },
    paper: {
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary,

        width: '100%',
        height: '80vh',
        display: 'flex',
        alignItems: 'center',
        flexDirection: 'column',
        minWidth: '600px'
      },

      messagesBody: {
        width: 'calc( 100% - 20px )', 
        margin: 10,
        overflowY: 'scroll', 
        height: 'calc( 100% - 80px )'
        },

        record: {
            width: '95%',
            height: 60,
            display: 'flex',
            alignItems: 'center',
            position: 'absolute',
            bottom: 0,
            left: 0,
            right: 0,
            margin: 10
        },
        message: {
            padding: 10,
            display: 'flex',
            justifyContent: 'space-between'
        }
  }));      

  const messages = ["lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol","lmao", "lol"]

function MainContent({ws}) {
    const classes = useStyles();

    const messageEndRef = useRef(null)
  
    const scrollToBottom = () => {
      messageEndRef.current.scrollIntoView({ behavior: "auto" })
    }
  
	useEffect(() => {
		let url = 'noop';
		(ws.logHistory === null) ?  url = `https://www.csunuav.me:1200/api/drone/logs/ssh/0/logs`: `https://www.csunuav.me:1200/api/drone/logs/ssh/${ws.logHistory.length}/logs`
		if(url !== 'noop') {		
	 		const fetchData = async () => {
				const result = await axios(url);
				setJsonData(result.data);
		    };
		    fetchData();
		}
    },[]);

    useEffect(scrollToBottom, [messages]);

    return (
        <main className={classes.content}>
            <div className={classes.toolbar} />
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Typography variant="h5" align='center'> Logs </Typography>
                        <Divider />
                            <div className={classes.messagesBody}>
                            {
                                    ws.logHistory.map(el =>
                                        (
                                            <div className={classes.message}>
                                             {el}
                                            </div>
                                        )
                                    )
                                }
                                <div ref={messageEndRef} />
                            </div>
                        </Paper>

                </Grid>
            </Grid>

        </main>
    )
}


export default MainContent;