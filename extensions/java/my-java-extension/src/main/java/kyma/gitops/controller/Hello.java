package kyma.gitops.controller;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class Hello {

    @RequestMapping("/")
    public String reply() {
        return "This is how you do Gitops with Kyma!!";
    }
}
