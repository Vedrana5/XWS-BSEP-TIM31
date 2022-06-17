package com.example.AgentApp.security;


import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.security.TokenUtilss;

import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class TokenAuthenticationFilter extends BasicAuthenticationFilter {

    private TokenUtilss tokenUtilss;
    private UserDetailsService userDetailsService;

    @Autowired
    private UserRepository userService;

    public TokenAuthenticationFilter(TokenUtilss tokenHelper, AuthenticationManager authManager, UserDetailsService userDetailsService) {
        super(authManager);
        this.tokenUtilss= tokenHelper;
        this.userDetailsService = userDetailsService;
    }

    @Override
    protected void doFilterInternal(HttpServletRequest req,
                                    HttpServletResponse res,
                                    FilterChain chain) throws IOException, ServletException {
        String header = req.getHeader("Authorization");

        if (header == null || !header.startsWith("Bearer ")) {
            chain.doFilter(req, res);
            return;
        }

        UsernamePasswordAuthenticationToken authentication = getAuthentication(req);

        SecurityContextHolder.getContext().setAuthentication(authentication);
        chain.doFilter(req, res);
    }

    private UsernamePasswordAuthenticationToken getAuthentication(HttpServletRequest request) {
        String token = tokenUtilss.getToken(request);
        System.out.println("TOKEN " + token);
        if (token != null) {
            // parse the token.
            String email = tokenUtilss.getEmailFromToken(token);
            System.out.println(email);


            if (email != null) {

                UserDetails userDetails = userDetailsService.loadUserByUsername(email);

                return new UsernamePasswordAuthenticationToken(email, null, userDetails.getAuthorities());
            }

            return null;
        }
        return null;
    }
}
