import { Body, Controller, Get, HttpCode, Post, UseGuards, UsePipes, ValidationPipe } from '@nestjs/common';
import { ApiBearerAuth, ApiOperation } from '@nestjs/swagger';
import { UserLoginDto } from './auth/user.login.dto';
import { AuthService } from './auth/auth.service';
import { RescanDto } from './dto/RescanDto';
import { RescanService } from './services/RescanService';
import { AuthGuard } from '@nestjs/passport';
import { GetSupportedChainsService } from './services/GetSupportedChainsService';

@Controller()
export class AppController {
  constructor(
    private readonly authService: AuthService,
    private readonly rescanService: RescanService,
    private readonly getSupportedChainsService: GetSupportedChainsService,
  ) {}

  @Get()
  @HttpCode(200)
  main(): boolean {
    return true;
  }

  @Get('/jobs')
  @HttpCode(200)
  jobs(): boolean {
    return true;
  }

  @Post('login')
  @ApiOperation({
    summary: 'Api for auth user',
  })
  @UsePipes(new ValidationPipe({ transform: true }))
  login(@Body() userLoginDto: UserLoginDto) {
    return this.authService.login(userLoginDto.login, userLoginDto.password);
  }

  @Post('/rescan')
  @ApiOperation({
    summary: 'Api for rescan',
  })
  @UsePipes(new ValidationPipe({ transform: true }))
  @ApiBearerAuth()
  @UseGuards(AuthGuard())
  rescan(@Body() dto: RescanDto) {
    return this.rescanService.rescan(dto.chainId, dto.from, dto.to);
  }

  @Get('/chains')
  @ApiOperation({
    summary: 'Api for getting supported chains',
  })
  getSupportedChains() {
    return this.getSupportedChainsService.get();
  }
}
