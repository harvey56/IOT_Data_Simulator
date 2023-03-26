import React from 'react';
import logo from './logo.svg';
import './App.css';
import Wrapper from './Wrapper';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import { ThemeContext } from '@emotion/react';
import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';

function App() {
  return (
    <Wrapper>
      <Box
        component="main"
        sx={{
          backgroundColor: (theme) =>
            theme.palette.mode === 'light'
              ? theme.palette.grey[100]
              : theme.palette.grey[900],
          flexGrow: 1,
          height: '100vh',
          overflow: 'auto',
          pl: 2,
          pr: 2,
        }}
      >
        <Toolbar />
        <Grid container spacing={2} sx={{ mb: 2 }}>
          <Grid item container xs={12} md={8} lg={9}>
            <Grid item xs={12} sx={{ mb: 2 }}>
              <Paper
                sx={{
                  p: 2,
                  display: 'flex',
                  flexDirection: 'column',
                  height: 240,
                  elevation: 2,
                }}
              >
                <PaperWrapper name="Map" />
              </Paper>
            </Grid>
            <Grid item xs={12} sx={{ mb: 2 }}>
              <Paper
                sx={{
                  p: 2,
                  display: 'flex',
                  flexDirection: 'column',
                  height: 240,
                  elevation: 2,
                }}
              ></Paper>
            </Grid>
          </Grid>
          <Grid item container sx={{ height: '100vh' }} xs={12} md={4} lg={3}>
            <Grid item>
              <Paper
                sx={{
                  p: 2,
                  display: 'flex',
                  flexDirection: 'column',
                  elevation: 2,
                }}
              >
                <PaperWrapper name="Device configuration" />
                <Box>
                  <Typography variant="body1" sx={{ fontWeight: 'bold' }}>
                    Configure MQTT
                  </Typography>
                  <Box marginBottom={1} marginTop={1}>
                    <Grid item container>
                      <Grid item xs={12} md={6}>
                        <FormControl
                          variant="standard"
                          sx={{ minWidth: 120, pr: 2 }}
                          fullWidth
                        >
                          <InputLabel id="demo-simple-select-standard-label">
                            Protocol
                          </InputLabel>
                          <Select
                            labelId="demo-simple-select-standard-label"
                            id="demo-simple-select-standard"
                            // value={age}
                            // onChange={handleChange}
                            label="Protocol"
                          >
                            <MenuItem value="">
                              <em>None</em>
                            </MenuItem>
                            <MenuItem value={10}>MQTT</MenuItem>
                            <MenuItem value={20}>WS</MenuItem>
                          </Select>
                        </FormControl>
                      </Grid>
                      <Grid item xs={12} md={6}>
                        <FormControl
                          variant="standard"
                          sx={{ minWidth: 120 }}
                          fullWidth
                        >
                          <InputLabel id="demo-simple-select-standard-label">
                            Port
                          </InputLabel>
                          <Select
                            labelId="demo-simple-select-standard-label"
                            id="demo-simple-select-standard"
                            // value={age}
                            // onChange={handleChange}
                            label="Port"
                          >
                            <MenuItem value="">
                              <em>None</em>
                            </MenuItem>
                            <MenuItem value={10}>1883</MenuItem>
                            <MenuItem value={20}>1884</MenuItem>
                            <MenuItem value={30}>8883</MenuItem>
                          </Select>
                        </FormControl>
                      </Grid>
                    </Grid>
                    <TextField
                      id="standard-basic"
                      label="Host"
                      variant="standard"
                      fullWidth
                    />
                  </Box>
                </Box>
                <Box>
                  <Typography variant="body1" sx={{ fontWeight: 'bold' }}>
                    Connection Status
                  </Typography>
                  <Box marginBottom={1} marginTop={1}>
                    <Box display={'flex'} justifyContent="flex-start">
                      <div
                        style={{
                          margin: '8px',
                          border: 'none',
                          borderRadius: '50px',
                          width: '24px',
                          height: '24px',
                          backgroundColor: 'green',
                        }}
                      ></div>
                      <Typography
                        variant="body2"
                        display={'flex'}
                        alignItems="center"
                      >
                        Connected
                      </Typography>
                    </Box>
                  </Box>
                </Box>
                <Typography variant="body1" sx={{ fontWeight: 'bold' }}>
                  Test MQTT Connection
                </Typography>
                <Box marginBottom={1} marginTop={1}>
                  <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                    Publish
                  </Typography>
                  <TextField
                    id="standard-basic"
                    label="Topic"
                    variant="standard"
                    fullWidth
                  />
                  <TextField
                    id="standard-basic"
                    label="QoS"
                    variant="standard"
                    fullWidth
                  />
                  <TextField
                    id="standard-basic"
                    label="Message"
                    variant="standard"
                    fullWidth
                  />
                  <Box margin={1}>
                    <Button size="small">Publish</Button>
                  </Box>
                  <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                    Subscribe
                  </Typography>
                  <TextField
                    id="standard-basic"
                    label="Topic"
                    variant="standard"
                    fullWidth
                  />
                  <Box margin={1}>
                    <Button size="small">Subscribe</Button>
                  </Box>
                </Box>
              </Paper>
            </Grid>
            <Grid item>
              <Paper
                sx={{
                  p: 2,
                  display: 'flex',
                  flexDirection: 'column',
                  elevation: 2,
                }}
              ></Paper>
            </Grid>
          </Grid>

          {/* <Grid container sx={{ mt: 4, mb: 4, ml: 4, mr: 4 }}>
            <Grid item xs={12}>
              <Paper
                sx={{
                  p: 2,
                  display: 'flex',
                  flexDirection: 'column',
                  elevation: 2,
                }}
              ></Paper>
            </Grid>
          </Grid> */}
        </Grid>
      </Box>
    </Wrapper>
  );
}

export function PaperWrapper({ name }: { name: string }) {
  return (
    <Box display={'flex'} justifyContent="space-between" marginBottom={2}>
      <div>
        <Typography variant="body1" sx={{ fontWeight: 'bold' }}>
          {name}
        </Typography>
      </div>
      <div>icon</div>
    </Box>
  );
}

export default App;
