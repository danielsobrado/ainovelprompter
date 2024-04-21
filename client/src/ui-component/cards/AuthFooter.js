// material-ui
import { Link, Typography, Stack } from '@mui/material';

// ==============================|| FOOTER - AUTHENTICATION 2 & 3 ||============================== //

const AuthFooter = () => (
  <Stack direction="row" justifyContent="space-between">
    <Typography variant="subtitle2" component={Link} href="https://drunsiel.com" target="_blank" underline="hover">
    drunsiel.com
    </Typography>
    <Typography variant="subtitle2" component={Link} href="https://danielsobrado.com" target="_blank" underline="hover">
      &copy; danielsobrado.com
    </Typography>
  </Stack>
);

export default AuthFooter;
